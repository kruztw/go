// ref: https://peterhpchen.github.io/2020/03/08/goroutine-and-channel.html

package main

import "fmt"

func deadlock() {
	ch := make(chan string) // should be `ch := make(chan string, 1)` or use go routine
	ch <- "FINISH"
}

func main() {
	//deadlock()

	ch := make(chan string)
	go func() {
		ch <- "FINISH"
	}()

	a := <-ch
	fmt.Printf("a = %v\n", a)
}
