// 參考: https://peterhpchen.github.io/2020/03/08/goroutine-and-channel.html
// <-ch : 從 channel 取值
// ch<- : 放值進 channel

package main

import (
    "fmt"
    "time"
)

func main() {
    total := 0
    ch := make(chan int, 1)
    ch <- total
    for i := 0; i < 1000; i++ {
        go func() {
            ch <- <-ch + 1
        }()
    }
    time.Sleep(time.Second)
    fmt.Println(<-ch)
}
