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

	if err = tv.AuthoriseClientKey("6c7b2ec679ffd1c2736abd621153eabb"); err != nil {
		log.Fatalf("could not authoise using client key: %v", err)
	}

	fmt.Println(tv.CurrentChannel())
	fmt.Println(tv.CurrentProgram())
}
