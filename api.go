package webos

import (
	"github.com/pkg/errors"
)

// MessageType is the type sent to and returned by the TV in the `type` field.
type MessageType string

// ErrorMessageType is returned by the TV when an error has occurred.
const ErrorMessageType MessageType = "error"

// RegisterMessageType is sent to the TV in a registration request.
const RegisterMessageType MessageType = "register"

// RegisteredMessageType is returned by the TV in response to a registration request.
const RegisteredMessageType MessageType = "registered"

// RequestMessageType is sent to the TV when issuing Commands.
const RequestMessageType MessageType = "request"

// ResponseMessageType is returned by the TV in response to a request.
const ResponseMessageType MessageType = "response"

// Message represents the JSON message format used in request and responses to
// and from the TV.
type Message struct {
	Type    MessageType `json:"type,omitempty"`
	ID      string      `json:"id,omitempty"`
	URI     Command     `json:"uri,omitempty"`
	Payload Payload     `json:"payload,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// Validate validates the Message.
//  Only used for response (type: response || registered) types.
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
		return errors.Errorf("unexpected API response type: %s", m.Type)
	}
}

// Payload represents the Payload contained in the Message body.
type Payload map[string]interface{}

// Validate valides the Payload.
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

// App represents an applications in the TVs responses.
type App struct {
	ReturnValue bool
	AppID       string
	WindowID    string
	ProcessID   string
	Running     bool
	Visible     bool
}

// Service represents services in the TVs responses.
type Service struct {
	Name    string
	Version float32
}

// ServiceList represents an array of Service types in the TVs responses.
type ServiceList struct {
	Services []Service
}

// Volume represents the audio output volume in the TVs responses.
type Volume struct {
	ReturnValue bool
	Scenario    string
	Volume      int32
	Muted       bool
}
