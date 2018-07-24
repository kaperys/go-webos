package webos

import (
	"github.com/pkg/errors"
)

type MessageType string

const ErrorMessageType MessageType = "error"
const RegisterMessageType MessageType = "register"
const RegisteredMessageType MessageType = "registered"
const RequestMessageType MessageType = "request"
const ResponseMessageType MessageType = "response"

type Message struct {
	Type    MessageType `json:"type,omitempty"`
	ID      string      `json:"id,omitempty"`
	URI     Command     `json:"uri,omitempty"`
	Payload Payload     `json:"payload,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func (m Message) Validate() error {
	switch m.Type {
	case ErrorMessageType:
		var err error
		if _, ok := m.Payload["returnValue"]; ok {
			err = m.Payload.Validate()
		}

		if err == nil {
			return errors.New(m.Error)
		}

		return errors.Errorf("API error: %s, %s", m.Error, err)
	case ResponseMessageType:
		return m.Payload.Validate()
	case RegisteredMessageType:
		if m.Payload == nil {
			return errors.New("empty payload")
		}
		return nil
	default:
		return errors.Errorf("unexpeced API response type: %s", m.Type)
	}
}

type Payload map[string]interface{}

func (p Payload) Validate() error {
	if p == nil {
		return errors.New("empty payload")
	}

	returnValueI, ok := p["returnValue"]
	if !ok {
		return errors.New("`returnValue` is missing")
	}

	returnValue, ok := returnValueI.(bool)
	if !ok {
		return errors.New("`returnValue` is not of type bool")
	}

	if !returnValue {
		if p["errorCode"] != nil {
			return errors.Errorf("error %v: %v", p["errorCode"], p["errorText"])
		}

		return errors.New("`returnValue` is false and `errorCode` is nil")
	}

	return nil
}
