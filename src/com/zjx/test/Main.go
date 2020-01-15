package main

import (
	"fmt"
	"unsafe"
)

func main() {
	testTypeSpace()
	//fmt.Println(testDefer())
}
/*
	测试defer的执行时机，return 返回并不是原子性的，分两步：
														返回值赋值，
														返回
	defer的执行时机介于二者之间，在defer函数中对返回值变量进行修改，不会影响返回的结果
 */
func testDefer() int {
	i := 10
	defer changeI(&i)
	return i
}

/*
	对于testDefer中的值进行改变
*/
func changeI(i *int) {
	*i += 10
}

/*
	测试数据类型所占的字节数
*/
func testTypeSpace() {
	var b byte
	fmt.Println(unsafe.Sizeof(b))
	var s string
	fmt.Println(unsafe.Sizeof(s))
	var i32 int32
	fmt.Println(unsafe.Sizeof(i32))
}
