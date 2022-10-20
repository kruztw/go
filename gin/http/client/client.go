package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type RegisterRequest struct {
	Msg string
}

func main() {

	body := RegisterRequest{Msg: "hello world"}

	bBody, err := json.Marshal(body)
	if err != nil {
		return
	}

	url := url.URL{Scheme: "http", Host: "127.0.0.1:8888", Path: "/post"}

	req, err := http.NewRequest(
		"POST",
		url.String(),
		bytes.NewReader(bBody),
	)

	if err != nil {
		return
	}

	req.Header.Add(
		"Authorization",
		"Bearer ",
	)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Failed to connect server: %v", err)
		return
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("resp.body = %v\n", string(bytes))
}
