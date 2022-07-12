// ref: https://colobu.com/2018/11/03/get-function-name-in-go/

package main

import (
	"fmt"
	"runtime"
)

func backtrace() {
	pc := make([]uintptr, 10) // at least 1 entry needed
	n := runtime.Callers(1, pc)
	frames := runtime.CallersFrames(pc[:n-2])
	for {
		frame, more := frames.Next()
		fmt.Printf("%s:%d %s\n", frame.File, frame.Line, frame.Function)
		if !more {
			break
		}
	}
}

func foo2() {
	pc1, _, _, _ := runtime.Caller(0)
	fmt.Printf("current: %v\n", runtime.FuncForPC(pc1).Name())
	pc2, _, _, _ := runtime.Caller(1)
	fmt.Printf("caller: %v\n", runtime.FuncForPC(pc2).Name())

	backtrace()
}

func foo1() {
	foo2()
}

func main() {
	foo1()
}
