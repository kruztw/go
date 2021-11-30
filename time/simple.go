package main

import (
    "fmt"
    "time"
)

func main() {
    for range time.Tick(time.Second*2) {    // 每 2s 觸發一次  
        time.Sleep(time.Second);            // 睡 1s
        fmt.Println(time.Now());            // 照理講應該是每 3s 印一次, 但事實上是每 2s 印一次
    }
}
