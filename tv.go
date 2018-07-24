package webos

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

// TV represents the TV. It contains the websocket connection, necessary channels
// used for communication and methods used for interaction with the TV.
type TV struct {
	ws      *websocket.Conn
	wsMutex sync.Mutex

	res      map[string]chan<- Message
	resMutex sync.Mutex
}

// NewTV dials the socket and returns a pointer to a new TV.
func NewTV(dialer *websocket.Dialer, addr string) (*TV, error) {
	ws, resp, err := dialer.Dial(fmt.Sprintf("wss://%s:3001", addr), nil)
	if err != nil {
		return nil, err
	}

	if err = resp.Body.Close(); err != nil {
		return nil, err
	}

	return &TV{ws: ws}, nil
}

// Command executes a Command on the TV.
func (tv *TV) Command(uri Command, req Payload) (Message, error) {
	return tv.request(&Message{
		Type:    RequestMessageType,
		ID:      requestID(),
		URI:     uri,
		Payload: req,
	})
}

// MessageHandler listens to the TVs websocket and reads responses.
// Responses are read into a Message type and added to appropriate channel
// based on the Message.ID.
func (tv *TV) MessageHandler() (err error) {
	defer func() {
		tv.resMutex.Lock()
		for _, ch := range tv.res {
			close(ch)
		}
		tv.res = nil
		tv.resMutex.Unlock()
	}()

	for {
		mt, p, err := tv.ws.ReadMessage()
		if err != nil {
			return err
		}

		if mt != websocket.TextMessage {
			continue
		}

		msg := Message{}

		err = json.Unmarshal(p, &msg)
		if err != nil {
			continue
		}

		tv.resMutex.Lock()
		ch := tv.res[msg.ID]
		tv.resMutex.Unlock()

		ch <- msg
	}
}

// AuthoriseClientKey autorises with the TV using an existing client key.
func (tv *TV) AuthoriseClientKey(key string) error {
	msg := Message{
		Type:    RegisterMessageType,
		ID:      requestID(),
		Payload: Payload{"client-key": key},
	}

	res, err := tv.request(&msg)
	if err != nil {
		return fmt.Errorf("could not make request: %v", err)
	}

	if rt := res.Type; rt != RegisteredMessageType {
		return fmt.Errorf("unexpected response type: %s", rt)
	}

	return nil
}

// AuthorisePrompt autorises with the TV using the PROMPT method.
func (tv *TV) AuthorisePrompt() (string, error) {
	msg := Message{
		Type:    RegisterMessageType,
		ID:      requestID(),
		Payload: pairPrompt(),
	}

	res, err := tv.request(&msg)
	if err != nil {
		return "", fmt.Errorf("could not make request: %v", err)
	}

	if rt := res.Type; rt != RegisteredMessageType {
		return "", fmt.Errorf("unexpected response type: %s", rt)
	}

	key := ""
	if k, ok := res.Payload["client-key"]; ok {
		k, ok := k.(string)
		if !ok {
			return "", errors.New("invalid client-key")
		}
		key = k
	}

	return key, nil
}

// Close closes the websocket connection to the TV.
func (tv *TV) Close() error {
	return tv.ws.Close()
}

// request makes a request to TV. It ensures a channel is available for responses
// using the given Message.ID and makes the request. Responses from the TV are added
// to the channel in the MessageHandler method, and read in this method. Responses
// are vaildates before they are returned.
func (tv *TV) request(msg *Message) (Message, error) {
	ch := make(chan Message, 1)
	tv.setupResponseChannel(msg.ID, ch)
	defer tv.teardownResponseChannel(msg.ID)

	b, err := json.Marshal(msg)
	if err != nil {
		return Message{}, fmt.Errorf("could not marshall request: %v", err)
	}

	tv.wsMutex.Lock()
	err = tv.ws.WriteMessage(websocket.TextMessage, b)
	tv.wsMutex.Unlock()

	if err != nil {
		return Message{}, fmt.Errorf("could not write to socket: %v", err)
	}

	for {
		select {
		case res, ok := <-ch:
			if !ok {
				return Message{}, errors.New("no response")
			}

			if res.Type == ResponseMessageType && msg.Type == RegisterMessageType {
				break
			}

			return res, res.Validate()
		case <-time.After(time.Second * 15):
			return Message{}, errors.New("timeout")
		}
	}
}

// setupResponseChannel ensures a channel is available for the given Message ID responses.
func (tv *TV) setupResponseChannel(id string, ch chan<- Message) {
	tv.resMutex.Lock()
	defer tv.resMutex.Unlock()

	if tv.res == nil {
		tv.res = make(map[string]chan<- Message)
	}

	tv.res[id] = ch
}

// teardownResponseChannel removes the channels used by the given Message ID.
func (tv *TV) teardownResponseChannel(id string) {
	tv.resMutex.Lock()
	defer tv.resMutex.Unlock()

	if ch, ok := tv.res[id]; ok {
		close(ch)
		delete(tv.res, id)
	}
}
