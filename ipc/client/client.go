package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net"

	"github.com/natefinch/npipe"
)

type Request struct {
	Action string          `json:"action"`
	Data   json.RawMessage `json:"data"`
}

type regReq struct {
	ConnectKey string `json:"connectKey"`
}

type StringIpc struct {
	c *bufio.Scanner
	m net.Conn
	t []byte
	l int
}

func (s *StringIpc) tokenize(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.Index(data, s.t); i >= 0 {
		return i + s.l, data[0:i], nil
	}
	if atEOF {
		return len(data), data, nil
	}
	return 0, nil, nil
}

func NewStringIpc(m net.Conn, spliter []byte) (c *StringIpc) {
	scanner := bufio.NewScanner(m)
	if scanner == nil {
		fmt.Println("memory error")
		return
	}

	c = &StringIpc{
		m: m,
		c: scanner,
		t: spliter,
		l: len(spliter),
	}
	if c == nil {
		fmt.Println("memory error")
		return
	}

	scanner.Split(c.tokenize)

	return
}

func (s *StringIpc) Read() (data []byte, err error) {
	if !s.c.Scan() {
		err = s.c.Err()
		return nil, err
	}

	return s.c.Bytes(), nil
}

func (s *StringIpc) Write(data []byte) (err error) {
	dataWithToken := append(data, s.t...)

	_, err = s.m.Write(dataWithToken)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func Reg(token string) {
	conn, err := npipe.Dial(`\\.\pipe\pipe1`)
	if err != nil {
		fmt.Println("npipe dial failed: %v", err)
		return
	}

	defer conn.Close()
	c := NewStringIpc(conn, []byte("^^2552*1814^^"))
	regreq := regReq{
		ConnectKey: token,
	}

	rawReq, err := json.Marshal(regreq)
	req := Request{
		Action: "register",
		Data:   rawReq,
	}
	reqM, err := json.Marshal(req)
	c.Write(reqM)

	res, err := c.Read()
	fmt.Printf("res: %v\n", res)
}

func main() {
	Reg("abc")
}
