package main

import (
	"fmt"
	. "github.com/cosiner/flag"
)

type Args struct {
	Arg1 int `names:"-i" default:"1" usage:"int arg"`
	Arg2 string `names:"-ss" default:"Hello" usage:"string arg"`
}

func main() {
	var arg Args

	set := NewFlagSet(Flag{})
	set.ParseStruct(&arg)

	fmt.Printf("Arg1: %d\nArg2: %s\n\n", arg.Arg1, arg.Arg2)
}

