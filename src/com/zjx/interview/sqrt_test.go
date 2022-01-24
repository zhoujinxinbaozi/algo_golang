/**
 *  @Author: JinxinZhou
 *  @Date  : 2022/1/24 2:49 下午
 */

package interview

import (
	"fmt"
	"testing"
)
/*
	求float的开平方
 */
func TestSqrt(t *testing.T) {
	//input := float64(43.56)
	input := float64(0.04)

	var low, high float64
	// 输入 > 1   结果一定大于1 且越平方值越大
	// 输入 < 1   结果一定小于1 且越平方值越小
	if input > 1 {
		low = float64(1)
		high = input
	} else {
		low = input
		high = float64(1)
	}
	for low < high {
		mid := (low + high)/2
		fmt.Println("mid = ", mid, "mid*mid", mid * mid)
		if mid * mid == input {
			fmt.Println("result : ", mid)
			return
		} else if mid * mid < input {
			low = mid
		} else {
			high = mid
		}
	}
}


