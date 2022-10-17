package main

import "fmt"

type foo struct {
	a int
	b string
}

func main() {

	var test interface{}
	test = foo{
		a: 1,
		b: "a",
	}

	res, ok := test.(foo)
	if !ok {
		fmt.Println("transform failed")
	}

	fmt.Printf("res: %v\n", res)

	test = &foo{
		a: 1,
		b: "a",
	}

	res2, ok := test.(*foo)
	if !ok {
		fmt.Println("transform failed")
	}

	fmt.Printf("res: %v\n", res2)

}
