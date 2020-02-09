package main

import "fmt"
/**
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例 1:

输入: 121
输出: true
示例 2:

输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3:

输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。
进阶:

你能不将整数转为字符串来解决这个问题吗？

solution:
	进阶的解法，通过取余和除法的操作，隔离出第一位和最后一位，进行比较，在去掉高位和地位进行循环即可
 */
func main() {
	fmt.Println(isPalindrome(1000021))
}
func isPalindrome(x int) bool {
	//边界判断
	if x < 0 {
		return false
	}
	div := 1
	for {
		if x/div < 10{
			break
		}
		div *= 10
	}
	for{
		if x <= 0{
			break
		}
		left := x / div
		right := x % 10
		if left != right{
			return false
		}
		x = (x % div) / 10 // 去掉高位和最低位
		div /= 100
	}
	return true
}

