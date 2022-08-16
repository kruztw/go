package main

// comment for golang

/* // comment for C
#include <stdio.h>
#include <stdlib.h>

static char *echo(char* s) {
   printf("%s\n", s);
   return s;
}

// no space between c code and import C !!
*/
import "C"
import "unsafe"
import "fmt"

func main() {
	cs := C.CString("Hello World")
	res := C.echo(cs)
	fmt.Printf("res = %v\n", C.GoString(res))
	C.free(unsafe.Pointer(cs))
}
