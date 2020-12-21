package main

/*
#cgo CFLAGS: -I./src
#cgo LDFLAGS: -L./lib -lmylib
#include "mylib.h"
*/
import "C"
import "fmt"

func main() {
	fmt.Println("Before")
	C.testCall()
	fmt.Println("After")
}
