package main

import "fmt"

type Easy interface {
	Add() int
}

type Hard interface {
	Multi() int
}

type Math interface {
	Easy
	Hard
}

type Arithmetic struct {
	a int
	b int
}

func (arith *Arithmetic) Add() int {
	return arith.a + arith.b;
}

func (arith *Arithmetic) Multi() int {
	return arith.a * arith.b;
}

func callAddMath(math Math) {
	fmt.Println("Math: ", math.Add())
}

func callAddEasy(easy Easy) {
	fmt.Println("Easy: ", easy.Add())
}

func main() {
	arith := [...]Arithmetic{
		{a:1, b:1},
		{a:2, b:2},
	}

	callAddMath(&arith[0])
	callAddEasy(&arith[1])
}
