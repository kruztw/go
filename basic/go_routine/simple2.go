package main

import (
	"fmt"
	"time"
)

func foo() {
	for {
		fmt.Printf("hello world\n")
		time.Sleep(time.Second)
	}
}

func main() {
	quit := make(chan bool)

	go func() {
		go foo()
		for {
			select {
			case <-quit:
				return
			default:
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	quit <- true
}
