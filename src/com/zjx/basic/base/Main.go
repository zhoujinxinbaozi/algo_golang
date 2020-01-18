package main

import "fmt"

func f1(a ...interface{}){ // 接受的参数会转化为slice
	fmt.Printf("type:%T, value:%#v\n", a, a)
}


func main() {
	s := []interface{}{1,2,"3",4}
	f1(s) // 将s整体当成一个元素传递进去
	f1(s...) // 将s中的每个元素分别取出来传递到函数中
	// type:[]interface {}, value:[[1 2 3 4]]
	// type:[]interface {}, value:[1 2 3 4]
}
