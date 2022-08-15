package main

/*
#cgo CFLAGS: -I./src
#cgo LDFLAGS: -L./src -lmylib
#include "mylib.h"
*/
import "C"

func main() {
	C.foo()
}
