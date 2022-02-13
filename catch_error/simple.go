package main

import "fmt"
import "runtime"


func catchPanic(e *error) {
	if err := recover(); err != nil {
		pc, file, line, ok := runtime.Caller(1)
		name := runtime.FuncForPC(pc).Name()
		*e = fmt.Errorf("%v: Catch panic \"%v\" in %s:%d (ok: %v)", file, err, name, line, ok)
	}
}

func Div(a int, b int) (ans int, err error) {
	defer catchPanic(&err);

	ans = a/b
	return
}

func main() {
	fmt.Println(Div(1, 1))
	fmt.Println(Div(1, 0))
}
