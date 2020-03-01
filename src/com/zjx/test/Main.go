package main

import (
	"fmt"
	"math"
	"reflect"
	"sync"
	"unsafe"
)

var wg sync.Once

func main() {
	// 使用Once调用函数，无论调用多少回，系统 都只会执行一次
	//wg.Do(once)
	//wg.Do(once)
	//testTypeSpace()
	//fmt.Println(testDefer())
	//judgeMax(3, 5)
	var m map[string]interface{}
	m = make(map[string]interface{}, 100)
	m["1"] = "1"
	m["2"] = 2
	m["3"] = 5.6
	m["4"] = []interface{}{1,"2",[]int{1,2}}
	PrintJSON(m)
}

// 输出map
func PrintJSON(m map[string]interface{}) {
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float", int64(vv))
		case int:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		case nil:
			fmt.Println(k, "is nil", "null")
		case map[string]interface{}:
			fmt.Println(k, "is an map:")
			PrintJSON(vv)
		default:
			fmt.Println(k, "is of a type I don't know how to handle ", fmt.Sprintf("%T", v))
		}
	}
}



func judgeMax(x, y int) {
	maxFloat := math.Max(float64(x), float64(y))
	max := int64(maxFloat)
	fmt.Println(reflect.TypeOf(max))
	fmt.Println(max)
}
func once() {

	fmt.Println("execute ...")
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
