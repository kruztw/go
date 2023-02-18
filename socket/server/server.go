package server

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	"../common"
)

const (
	socketName = "/tmp/a.sock"
	delim      = "\xde\xad\xbe\xef"
)

func Run() (err error) {
	defer os.RemoveAll(socketName)

	uaddr, err := net.ResolveUnixAddr("unix", socketName)
	if err != nil {
		return
	}

	listener, err := net.ListenUnix("unix", uaddr)
	if err != nil {
		return
	}

	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		fmt.Printf("accept failed (%v)", err)
		return
	}

	ipc := common.NewIPC(conn)
	if ipc == nil {
		fmt.Println("memory error")
		return
	}

	server := &common.Instance{
		Conn: conn,
		IPC:  ipc,
	}

	defer server.Close()

	rawRequest, hdlErr := server.IPC.Read()
	if rawRequest == nil {
		if hdlErr != nil {
			fmt.Printf("channel closed unexpectedly (%v)\n", hdlErr)
		} else {
			fmt.Printf("channel closed by client\n")
		}
	}

	fmt.Printf("server recv: %v\n", string(rawRequest))

	resp := common.Response{
		Code: 200,
		Msg:  "hello from server",
	}

	data, err := json.Marshal(resp)
	if err != nil {
		return
	}

	if err = server.IPC.Write(data); err != nil {
		return
	}

	return
}
