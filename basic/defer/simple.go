package main

import "fmt"

func foo() {
    a := 0
    defer fmt.Printf("eval first1: a = %d\n", a)
    defer func() { fmt.Printf("eval until func exit1: a = %d\n", a) }()
    a ++
    defer fmt.Printf("eval first2: a = %d\n", a)
    defer func() { fmt.Printf("eval until func exit2: a = %d\n", a) }()
}

func main() {
    foo()
}
