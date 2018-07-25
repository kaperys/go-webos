# Go WebOS 📺

A small Go library for interaction with webOS enabled TVs. Tested on LG webOS TV UH668V (webOS version 05.30.20).

[![Go Report Card](https://goreportcard.com/badge/github.com/kaperys/go-webos)](https://goreportcard.com/report/github.com/kaperys/go-webos)

```go
dialer := websocket.Dialer{
    HandshakeTimeout: 10 * time.Second,
    // the TV uses a self-signed certificate
    TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    NetDial: (&net.Dialer{Timeout: time.Second * 5}).Dial,
}

tv, err := webos.NewTV(&dialer, "<tv-ipv4-address>")
if err != nil {
    log.Fatalf("could not dial TV: %v", err)
}
defer tv.Close()

// the MessageHandler must be started to read responses from the TV
go tv.MessageHandler()

// AuthorisePrompt shows the authorisation prompt on the TV screen
key, err := tv.AuthorisePrompt()
if err != nil {
    log.Fatalf("could not authorise using prompt: %v", err)
}

// the key returned can be used for future request to the TV using the 
// AuthoriseClientKey(<key>) method, instead of AuthorisePrompt()
fmt.Println("Client Key:", key)

// see commands.go for available methods
tv.Notification("📺👌")
```

See [examples](examples/) for usage.

🌟 Inspired by [lgtv.js](https://github.com/msloth/lgtv.js), [go-lgtv](https://github.com/dhickie/go-lgtv) and [webostv](https://github.com/snabb/webostv).
