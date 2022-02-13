package main

import "fmt"


func Add1(a int, b int) int {
	return a + b
}

func Add2(a int, b int) (ans int) {
	ans = a + b
	return
}

func main() {
	fmt.Println(Add1(1, 2))
	fmt.Println(Add2(1, 2))
}
