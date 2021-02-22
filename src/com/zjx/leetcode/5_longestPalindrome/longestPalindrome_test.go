package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
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

// TestLongestPalindrome1
func TestLongestPalindrome(t *testing.T) {
	fmt.Println(longestPalindrome("165621"))
}

// longestPalindrome1
// 165621 --> #1#6#5#6#2#1#
// 遍历 #1#6#5#6#2#1#  分别找到最大满足条件的最小和最大下标 截取相应的字符串 和全局的结果进行比较
// 去掉# 返回
func longestPalindrome(s string) string {
	buf := bytes.Buffer{}
	for _, v := range s {
		buf.WriteRune('#')
		buf.WriteRune(v)
	}
	buf.WriteRune('#')

	tmp := buf.String()
	length := len(tmp)
	var result string

	fmt.Println(tmp)

	for index, _ := range tmp {
		before := index - 1
		after := index + 1
		var splitLeft, splitRight int
		for before >= 0 && after < length {
			if tmp[before] != tmp[after] {
				break
			}
			splitLeft = before
			splitRight = after
			before--
			after++
		}
		if splitLeft != splitRight {
			result = getResult(result, tmp[splitLeft:splitRight+1])
		}
	}
	return strings.ReplaceAll(result, "#", "")
}

// getResult
func getResult(s1, s2 string) string {
	if len(s1) > len(s2) {
		return s1
	}
	return s2
}

// TestRune
func TestRune(t *testing.T) {
	str := "顺利123"
	r := []rune(str)

	for _, v := range r {
		str := fmt.Sprintf("%c", v)
		fmt.Println(str)
		fmt.Println([]rune(str)[0])
	}
}