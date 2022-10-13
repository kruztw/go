package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket" //这里使用的是 gorilla 的 websocket 库
)

func main() {
	upgrader := websocket.Upgrader{}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		connect, err := upgrader.Upgrade(writer, request, nil)
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
	})

	err := http.ListenAndServe("127.0.0.1:8888", nil)
	if nil != err {
		log.Println(err)
		return
	}
}

func tickWriter(connect *websocket.Conn) {
	for {
		err := connect.WriteMessage(websocket.TextMessage, []byte("from server to client"))
		if nil != err {
			log.Println(err)
			break
		}

		time.Sleep(time.Second)
	}
}
