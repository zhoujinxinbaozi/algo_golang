package main

import "fmt"

/*
	闭包的含义：闭包是一个函数值，它引用了函数体之外的变量。
*/
func main() {
	myAdd := add()
	for i := 1; i <= 100; i++ {
		fmt.Println(myAdd(i))
	}
}

/*
	使用闭包的方式进行求1到n的累加和
*/

func add() func(current int) int {
	sum := 0
	return func(current int) int {
		sum += current
		return sum
	}
}
