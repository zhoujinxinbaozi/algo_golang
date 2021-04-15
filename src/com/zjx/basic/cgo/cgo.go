/**
 *  @Author: JinxinZhou
 *  @Date  : 2021/4/7 11:37 上午
 */


package main

//void SayHello(_GoString_ s);
//#include "stdio.h"
import "C"

import (
"fmt"
)

func main() {
	C.SayHello("Hello, World\n")
	C.SayHello("123")
	//C.max(1, 2)
}

//export SayHello
func SayHello(s string) {
	fmt.Print(s)
}
