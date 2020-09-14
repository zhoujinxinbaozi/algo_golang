package _func

import (
	"fmt"
	"testing"
)

// add 两个数之和
// @param a 加数
// @param b 加数
// @return a + b
func add(a, b int) int {
	return a + b
}

// sub 两数之差
func sub(a, b int) int {
	return a - b
}

// option
type option func(int, int) int

// execute
func Execute(option option, a, b int) int {
	return option(a, b)
}

// TestFunc
func TestFunc(t *testing.T) {
	fmt.Println(Execute(add, 1, 2))
	fmt.Println(Execute(sub, 1, 2))
}

// No2NameMap 学号和名字的映射关系
// key 学号
// value 姓名
type No2NameMap map[string]string

// nameSlice  姓名列表
type nameSlice []string

// TestTypeMapSlice 测试map slice
func TestTypeMapSlice(t *testing.T) {
	var m No2NameMap
	m = make(No2NameMap, 1)
	m["1"] = "周金鑫"
	for i, v := range m {
		fmt.Println(i)
		fmt.Println(v)
	}

	var nameList nameSlice
	nameList = append(nameList, "zjx")
	nameList = append(nameList, "yyl")
	fmt.Println(nameList)
}
