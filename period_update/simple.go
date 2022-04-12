package main

import (
	"fmt"
	"time"
	"sync"
)


func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	ticker := time.NewTicker(time.Duration(1) * time.Second)
	go func(timer *time.Ticker) {
		cnt := 0
		for {
			select {
			case <-ticker.C:
				cnt += 1
				fmt.Println("Hello ", cnt)
				break
			}

			if cnt == 5 {
				wg.Done()
			}
		}
	}(ticker)

	wg.Wait()
}
