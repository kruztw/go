// 參考: https://peterhpchen.github.io/2020/03/08/goroutine-and-channel.html
// <-ch : 從 channel 取值
// ch<- : 放值進 channel
// 當 channel 填滿時會造成 push 方等待 => channel size == 0 Unbuffered Channel, else buffered channel

package main

import (
    "fmt"
)

func main() {
    c := make(chan int, 10)
    go func() {
        for i := 0; i < 10; i++ {
            c <- i
        }
        close(c) // 關閉 Channel
    }()
    for i := range c { // 在 close 後跳出迴圈
        fmt.Println(i)
    }
}
