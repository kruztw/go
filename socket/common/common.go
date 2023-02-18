package common

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
)

type IPC struct {
	conn    net.Conn
	scanner *bufio.Scanner
	delim   []byte
}

type Request struct {
	Msg string `json:"msg"`
}

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type IChannel interface {
	Read() (data []byte, err error)
	Write(data []byte) (err error)
}

type Instance struct {
	Conn net.Conn
	IPC  IChannel
}

const (
	SocketName = "/tmp/a.sock"
	Delim      = "\xde\xad\xbe\xef"
)

func (obj *Instance) Close() {
	obj.Conn.Close()
}

func Listen() (listener net.Listener, err error) {
	if err = os.RemoveAll(SocketName); err != nil {
		return
	}

	uaddr, err := net.ResolveUnixAddr("unix", SocketName)
	if err != nil {
		return
	}

	listener, err = net.ListenUnix("unix", uaddr)
	if err != nil {
		return
	}

	return
}

func NewIPC(conn net.Conn) (ipc *IPC) {
	scanner := bufio.NewScanner(conn)
	if scanner == nil {
		return
	}

	ipc = &IPC{
		conn:    conn,
		scanner: scanner,
		delim:   []byte(Delim),
	}

	scanner.Split(ipc.tokenize)

	return
}

func (s *IPC) tokenize(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := bytes.Index(data, s.delim); i >= 0 {
		return i + len(s.delim), data[0:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	return 0, nil, nil
}

func (ipc *IPC) Read() (data []byte, err error) {
	if !ipc.scanner.Scan() {
		return nil, ipc.scanner.Err()
	}

	return ipc.scanner.Bytes(), nil
}

func (ipc *IPC) Write(data []byte) (err error) {
	dataWithToken := append(data, ipc.delim...)
	if _, err = ipc.conn.Write(dataWithToken); err != nil {
		return fmt.Errorf("write failed: %v\n", err)
	}

	return
}
