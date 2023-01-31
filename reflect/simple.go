package main

import (
	"fmt"
	"reflect"
)

type foo struct {
	A int
	B float64
	C string
	D []int
}

func main() {
	f := foo{A:1, C:"str", D:[]int{1,2,3}}
	fields := reflect.TypeOf(&f).Elem()
	vals := reflect.ValueOf(&f).Elem()

	for i := 0; i < vals.NumField(); i++ {
		field := fields.Field(i)

		if field.Name == "A" {
			vals.Field(i).SetInt(2)
		}

		if field.Type.Kind() == reflect.Slice {
			v := vals.Field(i)
			v = reflect.Append(v, reflect.ValueOf(4))
			vals.Field(i).Set(v)
		}

		fmt.Printf("field: %v\n", field)
		fmt.Printf("filed.Nmae: %v\n", field.Name)
		fmt.Printf("vals.Field(i) = %v\n", vals.Field(i))
	}
}

