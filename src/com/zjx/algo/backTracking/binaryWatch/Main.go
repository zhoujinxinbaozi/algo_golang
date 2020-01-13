package main

import (
	"fmt"
	"strings"
)

/***
	problem :

	二进制手表顶部有 4 个 LED 代表小时（0-11），底部的 6 个 LED 代表分钟（0-59）。
	每个 LED 代表一个 0 或 1，最低位在右侧。
	给定一个非负整数 n 代表当前 LED 亮着的数量，返回所有可能的时间。

	solution:
	无关痛痒


 */
func main() {
	//s1 := readBinaryWatch(1)
	//fmt.Println(s1)
	i := 1
	fmt.Printf("i = %09d", i)
}

/**

 */

func readBinaryWatch(num int) []string {
	var res []string
	for i := 0; i < 12; i++ {
		for j := 0; j < 60; j++ {
			h := fmt.Sprintf("%b", i)
			m := fmt.Sprintf("%b", j)
			countOne := strings.Count(h, "1") + strings.Count(m, "1")
			if countOne == num {
				res = append(res, fmt.Sprintf("%d:%02d", i, j))
			}
		}
	}
	return res
}
