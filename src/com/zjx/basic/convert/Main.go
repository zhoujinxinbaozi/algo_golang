package main

import (
	"fmt"
	"strconv"
)
/**
	字符串与数字之间的转换
 */
func main() {
	str := "123"
	i, _ := strconv.Atoi(str)
	fmt.Printf("%#v\n", i)

	str = strconv.Itoa(i)
	fmt.Printf("%#v", str)
}
