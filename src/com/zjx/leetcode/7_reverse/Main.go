package main

import (
	"fmt"
	"math"
)
/**
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:

输入: 123
输出: 321
 示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21

solution:
	先将x转为正数，每次取末尾的数，对全局的result进行更新，在根据x的正负，与int32的最大最小进行比较，溢出返回0
 */
func main() {
	fmt.Println(reverse(-2147483648))
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MinInt32)
}

func reverse(x int) int {
	current := math.Abs(float64(x))
	var result float64
	result = 0
	for {
		if current == 0 {
			break
		}
		m := math.Abs(float64(int32(current) % 10))
		result = result*10 + m
		current = float64(int32(current)/10)
	}
	if x < 0{
		if 0 - result > math.MinInt32{
			return int(0-result)
		}else{
			return 0
		}
	}else{
		if result > math.MaxInt32{
			return 0
		}else{
			return int(result)
		}
	}
}
