package webos

import (
	"crypto/tls"
	"fmt"
	"net"
	"time"

	"github.com/gorilla/websocket"
)

type Input struct {
	ws *websocket.Conn
}

// NewInput dials the socket and returns a pointer to a new Input.
func NewInput(uri string) (*Input, error) {
	dialer := websocket.Dialer{
		HandshakeTimeout: 10 * time.Second,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		NetDial: (&net.Dialer{
			Timeout: time.Second * 5,
		}).Dial,
	}
	ws, resp, err := dialer.Dial(uri, nil)
	if err != nil {
		return nil, err
	}

	if err = resp.Body.Close(); err != nil {
		return nil, err
	}

	return &Input{ws: ws}, nil
}

// SendButton
func (input *Input) SendButton(name string) error {
	body := fmt.Sprintf("type:button\nname:%s\n\n", name)
	err := input.ws.WriteMessage(websocket.TextMessage, []byte(body))
	if err != nil {
		return fmt.Errorf("could not write to socket: %v", err)
	}
	return nil
}

// Close closes the websocket connection.
func (input *Input) Close() error {
	return input.ws.Close()
}
