package main

import (
	"fmt"
	"sort"
	"reflect"
)

func main() {
	a := []int{1, 2}
	b := []int{2, 1}

	sort.Ints(a)
	sort.Ints(b)

	fmt.Println(reflect.DeepEqual(a, b))
}
