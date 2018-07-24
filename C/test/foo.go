package main

/*
 #cgo LDFLAGS: -L ./ -lfoo
 #include <stdlib.h>
 #include "foo.h"
*/
import "C"
import "fmt"

func main() {
	fmt.Println(C.count)
	C.foo()
}