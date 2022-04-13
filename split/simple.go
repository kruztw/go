package main

import (
	"fmt"
	"strings"
)

func main() {
	msg := "hello world"
	result := strings.Split(msg, " ")
	fmt.Println(result[0], result[1], len(result))
}
