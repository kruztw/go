package main

import "fmt"

func main() {
	data1 := []int{2, 3}
	data2 := append(data1, 4)
	fmt.Printf("data2: %v\n", data2)

	data3 := append([]int{1}, data1...)
	fmt.Printf("data3: %v\n", data3)
}
