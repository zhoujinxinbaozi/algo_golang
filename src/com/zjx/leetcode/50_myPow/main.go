package main

import "fmt"

/**
 * @Auther: zhoujinxin@bytedance.com
 * @Date: 2020/6/17 10:19
 */

/**
实现 pow(x, n) ，即计算 x 的 n 次幂函数。

示例 1:

输入: 2.00000, 10
输出: 1024.00000
示例 2:

输入: 2.10000, 3
输出: 9.26100
示例 3:

输入: 2.00000, -2
输出: 0.25000
解释: 2-2 = 1/22 = 1/4 = 0.25
*/
func main() {
	result := myPow(float64(2), 15)
	fmt.Println(result)
}

//func myPow(x float64, n int) float64 {
//	if n >= 0 {
//		return quickMul(x, n)
//	}
//	return 1.0 / quickMul(x, -n)
//}
//
//func quickMul(x float64, N int) float64 {
//	ans := 1.0
//	// 贡献的初始值为 x
//	x_contribute := x
//	// 在对 N 进行二进制拆分的同时计算答案
//	for N > 0 {
//		if N%2 == 1 {
//			// 如果 N 二进制表示的最低位为 1，那么需要计入贡献
//			ans *= x_contribute
//		}
//		// 将贡献不断地平方
//		x_contribute *= x_contribute
//		// 舍弃 N 二进制表示的最低位，这样我们每次只要判断最低位即可
//		N /= 2
//	}
//	return ans
//}

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quickMul1(x, n)
	}
	return 1.0 / quickMul1(x, -n)
}

// 递归
func quickMul(x float64, N int) float64 {
	if N == 0 {
		return 1
	}
	if N == 1 {
		return x
	}

	tmp := quickMul(x, N/2)
	if N&1 == 1 { //奇数
		return tmp * tmp * x
	}
	return tmp * tmp
}

// 非递归
func quickMul1(x float64, N int) float64 {
	if N == 0 {
		return 1
	}
	if N == 1 {
		return x
	}

	var tmp float64
	tmp = x
	result := 1.0

	for N != 0 {
		if N&1 == 1 { // 二进制最后一位为1 ，计入到结果中
			result *= tmp
		}
		tmp *= tmp
		N /= 2
	}

	return result

}
