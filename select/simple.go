// 參考: https://peterhpchen.github.io/2020/03/08/goroutine-and-channel.html

package main

import (
     "fmt"
     "time"
)

func main() {
    ch := make(chan string)

    go func() {
        fmt.Println("starts ...")
        time.Sleep(time.Second) // Heavy calculation
        fmt.Println("end")

        ch <- "FINISH"
    }()

    for {
        select {
        case <-ch: // Channel 中有資料執行此區域
            fmt.Println("received!!")
            return
        default: // Channel 阻塞的話執行此區域
            fmt.Println("WAITING...")
            time.Sleep(500 * time.Millisecond)
        }
    }
}
