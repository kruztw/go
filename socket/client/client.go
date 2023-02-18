package client

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"

	"../common"
)

func Run() (err error) {
	for {
		if _, err := os.Stat(common.SocketName); err == nil {
			break
		}

		time.Sleep(time.Second)
	}

	conn, err := net.Dial("unix", common.SocketName)
	if err != nil {
		return fmt.Errorf("dial faeild: %v", err)
	}

	ipc := common.NewIPC(conn)
	if ipc == nil {
		return fmt.Errorf("new string ipc failed: %v", err)
	}

	client := &common.Instance{
		Conn: conn,
		IPC:  ipc,
	}

	defer client.Close()

	req := common.Request{
		Msg: "Hello",
	}

	data, err := json.Marshal(req)
	if err != nil {
		return
	}

	if err := client.IPC.Write(data); err != nil {
		return fmt.Errorf("request failed: %v\n", err)
	}

	data, err = client.IPC.Read()
	if err != nil {
		return
	}

	fmt.Printf("client recv: %v\n", string(data))
	return nil
}
