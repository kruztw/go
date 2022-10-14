package main

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/natefinch/npipe"
)

type Request struct {
	Action string `json:"action"`
}

type Response struct {
	Data interface{} `json:"data"`
}

func main() {
	fmt.Printf("server start\n")

	listener, err := npipe.Listen(`\\.\pipe\pipe1`)
	if err != nil {
		fmt.Printf("npipe.Listen failed: %v\n", err)
		return
	}

	if err != nil {
		fmt.Printf("create ipc server failed: %v\n", err)
		return
	}

	conn, err := listener.Accept()
	if err != nil {
		fmt.Printf("accept failed (%v)\n", err)
	}

	defer conn.Close()

	for {
		rawRequest, err := readIPC(conn)
		if err != nil {
			fmt.Printf("readIPC failed: %v\n", err)
			break
		}

		fmt.Printf("server received: %v\n", string(rawRequest))
		handleRequest(conn, rawRequest)
	}

	err = listener.Close()
	if err != nil {
		fmt.Printf("ipc close failed: %v\n", err)
		return
	}

	return
}

func writeIPC(conn net.Conn, data []byte) error {
	n := byte(len(data))
	if n > 255 {
		return fmt.Errorf("data is too long")
	}

	msg := append([]byte{n}, data...)

	if _, err := conn.Write(msg); err != nil {
		return err
	}

	return nil
}

func readIPC(conn net.Conn) ([]byte, error) {
	len := make([]byte, 1)
	if _, err := conn.Read(len); err != nil {
		return nil, err
	}

	fmt.Printf("len = %v\n", len)

	body := make([]byte, len[0])
	n, err := conn.Read(body)
	if err != nil {
		return nil, err
	}

	fmt.Printf("body = %v\n", body)

	if n != int(len[0]) {
		return nil, fmt.Errorf("msg loss\n")
	}

	return body, nil
}

func handleRequest(conn net.Conn, rawRequest json.RawMessage) {
	var request Request
	resp := Response{
		Data: "HI",
	}

	defer func() {
		rawResponse, err := json.Marshal(resp)
		if err != nil {
			fmt.Printf("Encode response failed: %v\n", err)
			return
		}

		if err := writeIPC(conn, rawResponse); err != nil {
			fmt.Printf("Send response failed: %v\n", err)
			return
		}
	}()

	if err := json.Unmarshal(rawRequest, &request); err != nil {
		resp.Data = "Invalid"
		return
	}
}
