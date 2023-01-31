package main

import (
	"fmt"
	"reflect"
)

type foo struct {
	a int
	b float64
	c string
	d []int
}

func main() {
	f := foo{a:1, c:"str", d:[]int{1,2,3}}
	fields := reflect.TypeOf(&f).Elem()
	vals := reflect.ValueOf(&f).Elem()

	for i := 0; i < vals.NumField(); i++ {
		field := fields.Field(i)
		fmt.Printf("filed: %v\n", field)
		fmt.Printf("filed.Nmae: %v\n", field.Name)
		fmt.Printf("vals.Field(i) = %v\n", vals.Field(i))
	}
}

