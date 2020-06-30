package main

import (
	"fmt"
	"strings"
)

/**
 * @Auther: zhoujinxin@bytedance.com
 * @Date: 2020/6/20 21:41
 */

func main() {
	fmt.Println(lengthOfLastWord("a "))
}

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	arr := strings.Split(s, " ")
	if len(arr) == 0 {
		return 0
	}
	return len(arr[len(arr)-1])
}
