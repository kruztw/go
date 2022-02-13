package main

import "fmt"

type Math interface {
	Add() int
}

type Arithmetic struct {
	a int
	b int
}

func (arith *Arithmetic) Add() int {
	return arith.a + arith.b;
}

func callAdd(math Math) {
	fmt.Println("ADD: ", math.Add())
}


func main() {
	arith := Arithmetic{a:1, b:2}
	callAdd(&arith)
}
