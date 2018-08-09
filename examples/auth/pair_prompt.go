package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/gorilla/websocket"

	webos "github.com/kaperys/go-webos"
)

func main() {
	dialer := websocket.Dialer{
		HandshakeTimeout: 10 * time.Second,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		NetDial: (&net.Dialer{
			Timeout: time.Second * 5,
		}).Dial,
	}

	tv, err := webos.NewTV(&dialer, "192.168.1.67")
	if err != nil {
		log.Fatalf("could not dial: %v", err)
	}
	defer tv.Close()

	go tv.MessageHandler()

	key, err := tv.AuthorisePrompt()
	if err != nil {
		log.Fatalf("could not authorise using prompt: %v", err)
	}

	// this key can be used for future request to the TV using the AuthoriseClientKey method
	fmt.Println("Client Key:", key)

	tv.Notification("📺👌")
}
