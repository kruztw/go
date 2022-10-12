package main

import (
	"context"
	"fmt"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func main() {
	fmt.Println("Client started")

	conn, _, _, err := ws.DefaultDialer.Dial(context.Background(), "ws://127.0.0.1:8888/")
	if err != nil {
		fmt.Printf("connect failed: %v\n", err)
		return
	}

	fmt.Println("Connected to server")

	err = wsutil.WriteClientMessage(conn, ws.OpText, []byte("client hello"))
	if err != nil {
		fmt.Printf("send faileld: %v\n", err)
		return
	}

	bmsg, _, err := wsutil.ReadServerData(conn)
	if err != nil {
		fmt.Printf("receive failed: %v\n", err)
		return
	}

	fmt.Printf("client received: %v\n", string(bmsg))

	if err = conn.Close(); err != nil {
		fmt.Printf("close failed: %v\n", err)
		return
	}

	fmt.Println("Disconnected from server")
}
