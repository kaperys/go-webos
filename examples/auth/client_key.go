package main

import (
	"crypto/tls"
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

	if err = tv.AuthoriseClientKey("c219d8fbcee3839619dd80d6d9c57ad1"); err != nil {
		log.Fatalf("could not authorise using client key: %v", err)
	}

	tv.Notification("📺👌")
}
