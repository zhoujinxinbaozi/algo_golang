/**
 *  @Author: JinxinZhou
 *  @Date  : 2020/11/1 10:31 下午
 */
/**
实现函数double Power(double base, int exponent)，求base的exponent次方。不得使用库函数，同时不需要考虑大数问题。
 */
package _6_myPow

func myPow(x float64, n int) float64 {
	var nn int
	if n < 0 {
		nn = -n
	} else {
		nn = n
	}
	if n == 0 {
		return 1
	}
	var result float64
	tmp := myPow(x, nn/2)
	tmp *= tmp

	if nn & 1 == 1 {
		result = tmp * x
	} else {
		result = tmp
	}
	if n < 0 {
		return 1/result
	}
	return result
}