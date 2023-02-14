package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"syscall"
)

type stringIpc struct {
	c *bufio.Scanner
	m pipeInterface
	t []byte
	l int
}

type pipeInterface struct {
	f    *os.File
	path string
}

const (
	reqPipe = "/tmp/synoc2ia-ad-sync-service-req"
	resPipe = "/tmp/synoc2ia-ad-sync-service-res"
	delim   = "\xDE\xAD\xBE\xEF"
)

func (s *stringIpc) tokenize(data []byte, atEOF bool) (advance int, token []byte, err error) {
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

func newStringIpc(m pipeInterface, spliter []byte) (c *stringIpc, err error) {
	scanner := bufio.NewScanner(m)
	if scanner == nil {
		return nil, fmt.Errorf("memory error")
	}

	c = &stringIpc{
		m: m,
		c: scanner,
		t: spliter,
		l: len(spliter),
	}

	if c == nil {
		return nil, fmt.Errorf("memory error")
	}

	scanner.Split(c.tokenize)
	return
}

func createPipeInterface(path string) (pipeInterface, error) {
	pi := pipeInterface{
		path: path,
	}

	if err := pi.Open(); err != nil {
		return pi, err
	}

	return pi, nil
}

func (pipe *pipeInterface) Open() error {
	if pipe.f != nil {
		return nil
	}

	var err error
	if pipe.f, err = os.OpenFile(pipe.path, os.O_RDWR, os.ModeNamedPipe); err != nil {
		return fmt.Errorf("open file failed: %v", err)
	}

	return nil
}

func (pipe pipeInterface) Read(p []byte) (n int, err error) {
	return pipe.f.Read(p)
}

func (pipe *pipeInterface) Write(b []byte) (n int, err error) {
	if pipe.f == nil {
		return 0, fmt.Errorf("file is not open")
	}

	return pipe.f.Write(b)
}

func (pipe *pipeInterface) Close() {
	if pipe.f == nil {
		return
	}

	pipe.f.Close()
}

func (s *stringIpc) Read() (data []byte, err error) {
	if !s.c.Scan() {
		return nil, s.c.Err()
	}

	return s.c.Bytes(), nil
}

func (s *stringIpc) Write(data []byte) (err error) {
	dataWithToken := append(data, s.t...)

	if _, err = s.m.Write(dataWithToken); err != nil {
		return fmt.Errorf("string ipc write failed: %v", err)
	}

	return
}

func main() {

	os.Remove(reqPipe)
	os.Remove(resPipe)

	if err := syscall.Mkfifo(reqPipe, 0666); err != nil {
		fmt.Printf("mkfifo failed: %v\n", err)
		return
	}

	if err := syscall.Mkfifo(resPipe, 0666); err != nil {
		fmt.Printf("mkfifo failed: %v\n", err)
		return
	}

	piReq, err := createPipeInterface(reqPipe)
	if err != nil {
		fmt.Printf("create request named pipe interface failed: %v\n", err)
		return
	}

	piRes, err := createPipeInterface(resPipe)
	if err != nil {
		fmt.Printf("create response named pipe interface failed: %v\n", err)
		return
	}

	defer piReq.Close()
	defer piRes.Close()

	cReq, err := newStringIpc(piReq, []byte(delim))
	if err != nil {
		fmt.Printf("create request string ipc failed\n")
		return
	}

	cRes, err := newStringIpc(piRes, []byte(delim))
	if err != nil {
		fmt.Printf("create response string ipc failed\n")
		return
	}

	if err = cReq.Write([]byte("send request")); err != nil {
		fmt.Printf("channel write failed: %v\n", err)
		return
	}

	req, err := cReq.Read()
	if err != nil {
		fmt.Printf("channel read failed: %v\n", err)
		return
	}

	fmt.Printf("receive request: %v\n", string(req))

	if err = cRes.Write([]byte("send response")); err != nil {
		fmt.Printf("channel write failed: %v\n", err)
		return
	}

	resp, err := cRes.Read()
	if err != nil {
		fmt.Printf("channel read failed: %v\n", err)
		return
	}

	fmt.Printf("receive response: %v\n", string(resp))
}
