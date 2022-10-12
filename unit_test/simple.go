package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	ans := add(1, 2)
	fmt.Println(ans)
}
