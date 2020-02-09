package main

import (
	"fmt"
)

/***
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"

solution:
	遍历string，以当前的节点为中心，向两边扩散寻找最大(奇回文和偶回文需要注意，当前节点可以是单个或者是成对中的一个)
 */
func main() {
	fmt.Println(longestPalindrome("bb"))
}

func longestPalindrome(s string) string {
	if len(s) == 0{
		return s
	}
	result := string(s[0])
	for i := 0; i < len(s); i++ {
		current := string(s[i])
		left := i - 1
		right := i + 1
		var tmp string
		tmp = compare(left,right,current,result,s)
		if len(tmp) > len(result){
			result = tmp
		}
		current = ""
		tmp = compare(left+1,right,current,result,s)
		if len(tmp) > len(result){
			result = tmp
		}
	}
	return result
}

func compare(left, right int, current,result,s string)string{
	if left < 0 || right > len(s)-1{
		return ""
	}
	for {
		if left > -1 && right < len(s) {
			if s[left] == s[right] {
				current = current + string(s[right])
				current = string(s[left]) + current
				left--
				right++
				if len(current) > len(result) {
					result = current
				}
			} else {
				break
			}
		} else {
			break
		}
	}
	return result
}
