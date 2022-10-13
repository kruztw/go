package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket" //这里使用的是 gorilla 的 websocket 库
)

func main() {
	dialer := websocket.Dialer{}
	connect, _, err := dialer.Dial("ws://127.0.0.1:8888/", nil)
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
		case websocket.TextMessage: //文本数据
			fmt.Println(string(messageData))
		case websocket.BinaryMessage: //二进制数据
			fmt.Println(messageData)
		case websocket.CloseMessage: //关闭
		case websocket.PingMessage: //Ping
		case websocket.PongMessage: //Pong
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
