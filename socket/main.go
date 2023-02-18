package main

import (
	"fmt"
	"sync"

	"./client"
	"./server"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Add(1)

	go func() {
		if err := server.Run(); err != nil {
			fmt.Printf("server failed: %v\n", err)
		}

		wg.Done()
	}()

	go func() {
		if err := client.Run(); err != nil {
			fmt.Printf("client failed: %v\n", err)
		}

		wg.Done()
	}()

	wg.Wait()
}
