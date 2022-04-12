// 參考: https://peterhpchen.github.io/2020/03/08/goroutine-and-channel.html

package main

import (
    "fmt"
    "sync"
)

func foo (wg *sync.WaitGroup) {
    defer wg.Done();
    fmt.Println("hello from foo");
}

func main() {
    wg := new(sync.WaitGroup)
    wg.Add(1);                           // 等一個 thread 呼叫 Done
    go foo(wg);                          // 建立一個 goroutine

    wg.Wait();
}
