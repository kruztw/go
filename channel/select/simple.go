// ref: https://peterhpchen.github.io/2020/03/08/goroutine-and-channel.html

package main

import (
	"fmt"
	"time"
)

func main() {
    ch := make(chan string)

    go func() {
        fmt.Println("calculate goroutine starts calculating")
        time.Sleep(2*time.Second) // Heavy calculation
        fmt.Println("calculate goroutine ends calculating")

        ch <- "FINISH"
        time.Sleep(time.Second)
        fmt.Println("calculate goroutine finished")
    }()

    for {
        select {
        case <-ch: // Channel 中有資料執行此區域
            fmt.Println("main goroutine finished")
            return
        default: // Channel 阻塞的話執行此區域
            fmt.Println("WAITING...")
            time.Sleep(500 * time.Millisecond)
        }
    }
}
