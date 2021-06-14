package main

import (
    "fmt"
    "./foo"
)

func main() {
    foo.Hello()
    fmt.Println("Hello" + foo.World())
}
