package main

import (
	"fmt"
)

/**
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

solution:
	滑动窗口解决方案
*/

// main
func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

// lengthOfLongestSubstring 滑动窗口
func lengthOfLongestSubstring(s string) int {
	// key s[index] value is not used
	m := make(map[uint8]struct{}, len(s))
	// left point & right point
	var left, right int // 左指针代表从下标left开始 无重复的最大值
	// global result
	var result int
	// map contains s[index]  right++
	// map not contains s[index]  delete m[s[left]] left ++
	// every step update global result
	for right < len(s) {
		if _, ok := m[s[right]]; ok {
			delete(m, s[left]) // 从left到right 有重复的了 删除最左边的元素 在试试看
			left++
		} else {
			m[s[right]] = struct{}{}
			right++
		}
		result = Max(result, len(m))
	}
	return result
}

// Max
func Max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
