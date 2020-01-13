package main

import (
	"fmt"
	"unicode"
)

/**
	problem:
	给定一个字符串S，通过将字符串S中的每个字母转变大小写，我们可以获得一个新的字符串。返回所有可能得到的字符串集合。

	示例:
	输入: S = "a1b2"
	输出: ["a1b2", "a1B2", "A1b2", "A1B2"]

	输入: S = "3z4"
	输出: ["3z4", "3Z4"]

	输入: S = "12345"
	输出: ["12345"]

	solution:
	无关痛痒
 */
func main() {
	str := "a1b2"
	var slice []string
	slice = make([]string, 0)
	var target string
	solution(str, &slice, 0, target, 0)
	slice = removeReplicaSliceString(slice)
	fmt.Println(slice)
}

// 去除slice中的重复项
func removeReplicaSliceString(slice []string) []string {
	m1 := make(map[string]string, 10)
	result := make([]string, 0)
	for _, key := range slice {
		if _, ok := m1[key]; ok {
			continue
		} else {
			m1[key] = key
			result = append(result, m1[key])
		}
	}
	return result
}

// 回溯算法
func solution(str string, slice *[]string, count int, target string, start int) {
	if start == len(str) - 1 {
		*slice = append(*slice, target)
		target = ""
		return
	}

	for i := start; i < len(str); i++ {
		target = target + string(str[start])
		//fmt.Println("before dg : ", target)
		solution(str, slice, count+1, target, start+1)
		target = target[0 : len(target)-1]
		target = target + string(unicode.ToUpper(rune(str[start])))
		//fmt.Println("after dg : ", target)
		solution(str, slice, count+1, target, start+1)
		target = target[0 : len(target)-1]
	}

}
