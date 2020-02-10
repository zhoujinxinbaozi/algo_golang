package main

import (
	"fmt"
	"math"
)
/**
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。

solution:
	找出最短的字符串，从前向后，一次拿出每个字节，进行比较，相同加入到result中，返回最终的结果
 */
func main() {
	fmt.Println(longestCommonPrefix([]string{"a"}))
}

func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	}
	result := ""
	length := math.MaxInt64
	// 找到字符串中最短的长度
	for _, value := range strs {
		if len(value) < length {
			length = len(value)
		}
	}
	if length <= 0 {
		return result
	}
	current := 0
	for {
		if current > length-1 {
			break
		}
		tmp := strs[0][current]
		equals := true
		for _, value := range strs {
			if tmp != value[current] {
				equals = false
				break
			}
		}
		if equals {
			result = result + string(tmp)
			current++
		} else {
			break
		}
	}
	return result
}
