package main

import "fmt"

func main() {
    var a, b int = 1, 2
    c := 3
    fmt.Printf("%d + %d = %d\n", a, b, c)

    ary := make([]int, 3) // int array with 3 elements
    ary[0] = 1
    ary = append(ary, 4)

    fmt.Printf("ary[0] = %d, ary[3] = %d\n", ary[0], ary[3])

    var heap *int = new(int)
    *heap = 1
    fmt.Printf("[%p] = %d\n", heap, *heap)
}
