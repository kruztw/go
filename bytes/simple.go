package main

import (
	"fmt"
)

func main() {
	s := "abc"
	fmt.Println([]byte(s))

	b := []byte("abc")
	fmt.Println(string(b))
}
