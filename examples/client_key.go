package main

import (
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/websocket"

	webos "gitlab.com/kaperys/go-webos"
)

func main() {
	dialer := websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 10 * time.Second,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		NetDial: (&net.Dialer{
			Timeout:   time.Second * 5,
			KeepAlive: time.Second * 30,
		}).Dial,
	}

	tv, err := webos.NewTV(&dialer, "192.168.1.67")
	if err != nil {
		log.Fatalf("could not dial: %v", err)
	}
	defer tv.Close()

	go tv.MessageHandler()

	if err = tv.AuthoriseClientKey("c219d8fbcee3839619dd80d6d9c57ad1"); err != nil {
		log.Fatalf("could not authoise using client key: %v", err)
	}

	tv.Notification("📺👌")
}
