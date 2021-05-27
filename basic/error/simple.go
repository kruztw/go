package main

import (
    "fmt"
)

func main() {
    var err error
    if err != nil {
        fmt.Println(err.Error())
        panic(err)
    }
}
