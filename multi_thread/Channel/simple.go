// 參考: https://peterhpchen.github.io/2020/03/08/goroutine-and-channel.html

package main

import (
    "fmt"
)

func foo1 (c chan string) {
    fmt.Println("hello from foo1");
    c <- "finish1";
}

func foo2 (c chan string) {
    fmt.Println("hello from foo2");
    c <- "finish2";
}

func main() {
    ch := make(chan string)

    go foo1(ch);                          // 建立一個 goroutine
    go foo2(ch);

    <-ch;
    <-ch;
}
