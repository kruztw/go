package main

import "fmt"

func main() {
	var b [256]byte
	for i := 0; i < 256; i += 1 {
		b[i] = uint8(i)
	}

	fmt.Printf("b = %v\n", b)
}
