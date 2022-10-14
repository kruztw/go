package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"sync"

	"github.com/natefinch/npipe"
)

const (
	E_NONE = iota
	E_INTERNAL
	E_INVALID_REQUEST
	E_AUTH_FAIL           // e.g. SRP verification failed
	E_AUTH_REQUIREDED     // e.g. get_random_key should come after auth succeed
	E_NO_PERMISSION       // e.g. server response forbidden for user info request
	E_CONNECTION_ERROR    // e.g. unable to connect to server, connection timeout ...
	E_USER_STATUS_UNKNOWN // e.g. user status cannot be specified
	E_USER_DISABLED       // e.g. run SRP on disabled user
	E_USER_STAGED         // e.g. user is not yet activated on server
	E_USER_PWD_PENDING    // e.g. user is waiting for activation from mail
	E_USER_PWD_EXPIRED    // e.g. user password is reset on server
)

type IHandler interface {
	Handle(msgData json.RawMessage) (resp Response)
	AfterHandle(resp Response)
}

type Config struct {
	Handler IHandler
}

type Request struct {
	Action string          `json:"action"`
	Data   json.RawMessage `json:"data"`
}

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type ErrorResult struct {
	Code   int    `json:"code"`
	Reason string `json:"reason"`
}

func Addr() string {
	return `\\.\pipe\pipe1`
}

func Listen(address string) (listener net.Listener, err error) {
	return npipe.Listen(address)
}

func NewErrorResult(code int, reason string) ErrorResult {
	return ErrorResult{
		Code:   code,
		Reason: reason,
	}
}

func main() {
	StartIpcThread()
}

func (r *Response) SetData(success bool, data interface{}) (err error) {
	r.Success = success
	r.Data = data
	return
}

func StartIpcThread() {
	fmt.Printf("[ipcThread] start")

	var ipcServer *IpcServer = nil
	var lock sync.Mutex

	for {

		var err error

		lock.Lock()
		ipcServer, err = newIpcServer(Addr())
		lock.Unlock()

		if err != nil {
			fmt.Printf("create ipc server failed: ", err)
			return
		}

		for {
			err := ipcServer.Handle()
			if err != nil {
				break
			}
		}

		ipcServer.Close()
	}

	lock.Lock()
	if ipcServer != nil {
		ipcServer.Close()
	}
	lock.Unlock()

	fmt.Printf("[ipcThread] stopped")
}

type IChannel interface {
	Read() (data []byte, err error)
	Write(data []byte) (err error)
}

type IPipe interface {
	Read(in []byte) (out []byte, err error)
	Write(in []byte) (out []byte, err error)
}

type ChannelPipe struct {
	c IChannel
	d IPipe
}

type IpcServer struct {
	listener net.Listener
	stopped  bool
	abort    (chan bool)
	wg       sync.WaitGroup
	lock     sync.Mutex
}

type channelContext struct {
	authed      bool
	otpVerified bool
}

type StringIpc struct {
	c *bufio.Scanner
	m net.Conn
	t []byte
	l int
}

func newIpcServer(address string) (server *IpcServer, err error) {
	listener, err := Listen(address)
	if err != nil {
		return
	}

	return &IpcServer{
		listener: listener,
		stopped:  false,
		abort:    make(chan bool),
		lock:     sync.Mutex{},
	}, nil
}

func (s *IpcServer) Close() (err error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.stopped {
		return
	}

	s.stopped = true
	close(s.abort)
	err = s.listener.Close()
	if err != nil {
		fmt.Printf("ipc close failed %v", err)
	}
	s.wg.Wait()
	return err
}

func NewChannelPipe(c IChannel) (p *ChannelPipe) {
	return &ChannelPipe{
		c: c,
	}
}

func (s *IpcServer) Handle() (err error) {
	s.wg.Add(1)
	conn, err := s.listener.Accept()
	s.wg.Done()

	if err != nil {
		select {
		case <-s.abort:
			fmt.Printf("ipc need to stop, stop handle")
			return
		default:
			fmt.Printf("accept failed (%v)", err)
			return
		}
	}

	s.wg.Add(1)

	go func(conn net.Conn) {
		ctx := &channelContext{}

		defer func() {
			conn.Close()
			s.wg.Done()
		}()

		c := NewStringIpc(conn, []byte("^^2552*1814^^"))
		if c == nil {
			fmt.Printf("memory error")
			return
		}
		p := NewChannelPipe(c)
		if p == nil {
			fmt.Printf("memory error")
			return
		}

		for {
			rawRequest, hdlErr := p.Read()
			if rawRequest == nil {
				if hdlErr != nil {
					fmt.Printf("channel closed unexpectedly (%v)", hdlErr)
				} else {
					fmt.Printf("channel closed by client")
				}
				break
			}
			s.handleRequest(ctx, rawRequest, p)
		}
	}(conn)
	return
}

func (s *StringIpc) Write(data []byte) (err error) {
	dataWithToken := append(data, s.t...)

	_, err = s.m.Write(dataWithToken)
	if err != nil {
		fmt.Print(err)
		return
	}
	return
}

func (s *StringIpc) Read() (data []byte, err error) {
	if !s.c.Scan() {
		err = s.c.Err()
		return nil, err
	}

	return s.c.Bytes(), nil
}

func (p *ChannelPipe) Read() (out []byte, err error) {
	data, err := p.c.Read()
	if err != nil {
		return
	}
	if data == nil {
		return
	}

	if p.d != nil {
		data, err = p.d.Read(data)
		if err != nil {
			return
		}
	}

	return data, nil
}

func NewStringIpc(m net.Conn, spliter []byte) (c *StringIpc) {
	scanner := bufio.NewScanner(m)
	if scanner == nil {
		fmt.Printf("memory error")
		return
	}

	c = &StringIpc{
		m: m,
		c: scanner,
		t: spliter,
		l: len(spliter),
	}
	if c == nil {
		fmt.Printf("memory error")
		return
	}

	scanner.Split(c.tokenize)

	return
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

func (s *IpcServer) handleRequest(ctx *channelContext, rawRequest json.RawMessage, p *ChannelPipe) {
	var request Request
	errResult := ErrorResult{
		Code: E_NONE,
	}

	err := json.Unmarshal(rawRequest, &request)
	if err != nil {
		fmt.Printf("Fail to parse request (%v, %v)", err, string(rawRequest))
		errResult = NewErrorResult(E_INVALID_REQUEST, "invalid request")
		responseError(p, errResult)
		return
	}

	var config Config
	actionValid := false
	resp := Response{}

	defer func() {
		if actionValid {
			config.Handler.AfterHandle(resp)
		}
		if errResult.Code != E_NONE {
			responseError(p, errResult)
		}
	}()

	resp = Response{}
	rawResponse, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Encode response failed (%v)", err)
		return
	}

	err = p.Write(rawResponse)
	if err != nil {
		fmt.Printf("Send response failed (%v)", err)
		return
	}
}

func responseError(c IChannel, msg ErrorResult) {
	resp := Response{
		Success: false,
		Data:    msg,
	}
	rawResponse, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Encode response failed (%v)", err)
		return
	}
	err = c.Write(rawResponse)
	if err != nil {
		fmt.Printf("Send response failed (%v)", err)
		return
	}
}

func (p *ChannelPipe) Write(data []byte) (err error) {
	if p.d != nil {
		data, err = p.d.Write(data)
		if err != nil {
			return
		}
	}

	return p.c.Write(data)
}
