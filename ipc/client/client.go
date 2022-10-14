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

func readIPC(conn net.Conn) ([]byte, error) {
	len := make([]byte, 1)
	if _, err := conn.Read(len); err != nil {
		return nil, err
	}

	body := make([]byte, len[0])
	n, err := conn.Read(body)
	if err != nil {
		return nil, err
	}

	if n != int(len[0]) {
		return nil, fmt.Errorf("msg loss\n")
	}

	return body, nil
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

func main() {
	conn, err := npipe.Dial(`\\.\pipe\pipe1`)
	if err != nil {
		fmt.Printf("npipe dial failed: %v\n", err)
		return
	}

	defer conn.Close()

	req := Request{
		Action: "hello",
	}

	rawReq, err := json.Marshal(req)
	if err != nil {
		fmt.Printf("Marshal failed: %v\n", err)
		return
	}

	if err := writeIPC(conn, rawReq); err != nil {
		fmt.Printf("writeIPC failed: %v\n", err)
		return
	}

	resp, err := readIPC(conn)
	if err != nil {
		fmt.Printf("readIPC failed: %v\n", err)
		return
	}

	fmt.Printf("response: %v\n", string(resp))
}
