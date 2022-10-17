package main

import (
	"fmt"
	"time"
)

func main() {
	quit := make(chan bool)
	go func() {
		fmt.Printf("hello world\n")
		quit <- true
	}()

	for {
		select {
		case <-quit:
			return
		default:
			time.Sleep(time.Second)
		}
	}
}
