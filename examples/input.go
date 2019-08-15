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

	tv, err := webos.NewTV(&dialer, "192.168.1.3")
	if err != nil {
		log.Fatalf("could not dial: %v", err)
	}
	defer tv.Close()

	go tv.MessageHandler()

	if err = tv.AuthoriseClientKey("284d99ac14a106d1004557321dfd7d86"); err != nil {
		log.Fatalf("could not authoise using client key: %v", err)
	}

	tv.KeyDown()

	tv.KeyOk()
}
