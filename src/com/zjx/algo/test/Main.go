package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var b byte
	fmt.Println(unsafe.Sizeof(b))
	var s string
	fmt.Println(unsafe.Sizeof(s))
	var i32 int32
	fmt.Println(unsafe.Sizeof(i32))
}
