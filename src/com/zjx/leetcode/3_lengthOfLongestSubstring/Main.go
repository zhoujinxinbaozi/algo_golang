package main

import (
	"fmt"
	"strings"
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
	map存放 key=s[i] value=index, j 表示 当前元素向前，最先出现重复元素的下一个位置，result取max(result, i-j+1)
 */
func main() {
	lengthOfLongestSubstring("abba")
}

func lengthOfLongestSubstring(s string) int {
	var m map[int]int
	m = make(map[int]int, len(s))
	var j int // 表示的是上一次出现这个字符的位置    --> 当前元素向前，最先出现重复元素的下一个位置
	result := 0
	for index, value := range s {
		//fmt.Printf("index=%d, value=%v", index, value)
		if v, ok := m[int(value)]; ok {
			if v+1 > j {  // 在当前这一段区间里，最长的不重复子串应该是   max（当前位置上一次出现的位置之差，当前位置-j)
				      // abba 在遍历到最后一个a的时候，i=3，此时j为2，上一个出现a的位置为0
				j = v + 1
			}
		}
		m[int(value)] = index
		if result < (index - j + 1) {
			result = index - j + 1
		}
		fmt.Println(result)
	}
	return result
}
// 滑动窗口
func lengthOfLongestSubstring1(s string) int {
	var Length int
	var s1 string
	left := 0
	right := 0
	s1 = s[left:right]
	for ; right < len(s); right++ {

		if index := strings.IndexByte(s1, s[right]); index != -1 {
			left += index + 1
		}
		s1 = s[left : right+1]
		if len(s1) > Length {
			Length = len(s1)
		}
	}

	return Length
}
