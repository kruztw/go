package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type RegisterRequest struct{}

func main() {

	body := RegisterRequest{}

	bBody, err := json.Marshal(body)
	if err != nil {
		return
	}

	url := url.URL{Scheme: "https", Host: "127.0.0.1:8888", Path: "/"}

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

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Failed to connect server: %v", err)
		return
	}

	defer resp.Body.Close()

}
