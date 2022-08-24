package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	once sync.Once
)

func ReadConfig() {
	fmt.Println("call ReadConfig")
	once.Do(func() {
		fmt.Println("init config: only do once")
	})
}

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			ReadConfig()
		}()
	}
	time.Sleep(time.Second)
}
