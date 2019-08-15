package webos

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

var (
	// Protocol is the protocol used to connect to the TV.
	Protocol = "wss"

	// Port is the port used to connect to the TV.
	Port = 3001
)

// TV represents the TV. It contains the websocket connection, necessary channels
// used for communication and methods used for interaction with the TV.
type TV struct {
	ws      *websocket.Conn
	wsMutex sync.Mutex

	res      map[string]chan<- Message
	resMutex sync.Mutex
	input    *Input
}

// NewTV dials the socket and returns a pointer to a new TV.
func NewTV(dialer *websocket.Dialer, ip string) (*TV, error) {
	addr := fmt.Sprintf("%s://%s:%d", Protocol, ip, Port)
	ws, resp, err := dialer.Dial(addr, nil)
	if err != nil {
		return nil, err
	}

	if err = resp.Body.Close(); err != nil {
		return nil, err
	}
	tv := &TV{ws: ws}
	return tv, nil
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
	if tv.input != nil {
		tv.input.Close()
		tv.input = nil
	}
	return tv.ws.Close()
}

// request makes a request to TV. It ensures a channel is available for responses
// using the given Message.ID and makes the request. Responses from the TV are added
// to the channel in the MessageHandler method, and read in this method. Responses
// are vaildates before they are returned.
func (tv *TV) request(msg *Message) (Message, error) {
	ch := tv.setupResponseChannel(msg.ID)
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
func (tv *TV) setupResponseChannel(id string) chan Message {
	tv.resMutex.Lock()
	defer tv.resMutex.Unlock()

	if tv.res == nil {
		tv.res = make(map[string]chan<- Message)
	}

	ch := make(chan Message, 1)
	tv.res[id] = ch
	return ch
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

// requestID returns a random 8 character string. Requests and Responses sent to and from
// the TV are linked by this ID.
func requestID() string {
	rs := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, 8)
	for i := range b {
		b[i] = rs[rand.Intn(len(rs))]
	}
	return string(b)
}

// createInput create if needed an input
func (tv *TV) createInput() (*Input, error) {
	msg := Message{
		Type: RequestMessageType,
		ID:   requestID(),
		URI:  GetPointerInputSocketCommand,
	}
	res, err := tv.request(&msg)
	if err != nil {
		return nil, fmt.Errorf("could not make request: %v", err)
	}
	var socketPath string
	socketPath = fmt.Sprintf("%s", res.Payload["socketPath"])

	input, err := NewInput(socketPath)
	if err != nil {
		return nil, fmt.Errorf("could not dial: %v", err)
	}
	return input, nil
}
