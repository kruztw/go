package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func get() {
	resp, err := http.Get("http://localhost:8888/")
	if err != nil {
		fmt.Printf("http Get failed: %v", err)
		return
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("response: %v\n", string(content))
}

func post() {
	resp, err := http.PostForm("http://localhost:8888", url.Values{"key1": {"val1"}, "key2": {"val2"}})
	if err != nil {
		fmt.Printf("http post failed: %v", err)
		return
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("response: %v\n", string(content))
}

func main() {
	get()
	post()
}
