// 參考: https://peterhpchen.github.io/2020/03/08/goroutine-and-channel.html

package main

import (
    "fmt"
    "time"
)

func foo () {
    fmt.Println("hello from foo");
}

func main() {
    go foo();
    go func() {
        fmt.Println("hello from main");
    }()

    time.Sleep(1 * time.Second)
}
