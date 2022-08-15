package main

// comment for golang

// // comment for C
// #include <stdio.h>
// #include <stdlib.h>
//
// static void myprint(char* s) {
//   printf("%s\n", s);
// }
// // no space between c code and import C !!
import "C"
import "unsafe"

func main() {
	cs := C.CString("Hello World")
	C.myprint(cs)
	C.free(unsafe.Pointer(cs))
}
