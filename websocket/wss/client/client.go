package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	dialer := websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 45 * time.Second,
	}

	dialer.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	connect, _, err := dialer.Dial("wss://127.0.0.1:8888/", nil)
	if nil != err {
		log.Println(err)
		return
	}

	defer connect.Close()
	go tickWriter(connect)

	for {
		messageType, messageData, err := connect.ReadMessage()
		if nil != err {
			log.Println(err)
			break
		}
		switch messageType {
		case websocket.TextMessage:
			fmt.Println(string(messageData))
		case websocket.BinaryMessage:
			fmt.Println(messageData)
		case websocket.CloseMessage:
		case websocket.PingMessage:
		case websocket.PongMessage:
		default:

		}
	}
}

func tickWriter(connect *websocket.Conn) {
	for {
		err := connect.WriteMessage(websocket.TextMessage, []byte("from client to server"))
		if nil != err {
			log.Println(err)
			break
		}

		time.Sleep(time.Second)
	}
}
