package main

/*
#include <stdio.h>
#include <stdlib.h>

void myprint(char* s) {
	printf("%s\n", s);
}
*/
import "C"
import (
	"./myrouter"
	"unsafe"
)

func Example() {
	cs := C.CString("Hello from stdio\n")
	C.myprint(cs)
	C.free(unsafe.Pointer(cs))

	myrouter.My()
}

func main() {
	Example()
}
