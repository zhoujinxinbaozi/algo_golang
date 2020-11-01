/**
 *  @Author: JinxinZhou
 *  @Date  : 2020/11/1 8:26 上午
 */
/**
给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。
 */
package _57_reverseWords

import (
	"fmt"
	"testing"
)

func reverseWords(s string) string {
	var tmp []byte
	var result string
	for i, v := range s {
		if v == ' ' || i == len(s)-1 {
			if i == len(s) -1 {
				tmp = append(tmp, byte(v))
			}
			for t := len(tmp)-1; t >= 0; t-- {
				result += string(tmp[t])
			}
			tmp = nil
			if i != len(s)-1 {
				result += " "
			}
			continue
		}
		tmp = append(tmp, byte(v))
	}
	return result
}

func TestReverseWords(t *testing.T) {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}