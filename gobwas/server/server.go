package main

import (
	"fmt"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func main() {
	http.ListenAndServe(":8888", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Client connected")
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			fmt.Println("Error starting socket server: " + err.Error())
			return
		}

		go func() {
			defer conn.Close()
			for {
				bmsg, op, err := wsutil.ReadClientData(conn)
				if err != nil {
					fmt.Printf("receive failed: %v\ndisconnect\n", err)
					return
				}

				fmt.Printf("server received: %v\n", string(bmsg))

				if err := wsutil.WriteServerMessage(conn, op, []byte("server hello")); err != nil {
					fmt.Printf("send failed: %v\ndisconnect\n", err)
					return
				}
			}
		}()
	}))
}
