/**
 *  @Author: JinxinZhou
 *  @Date  : 2020/10/28 10:58 上午
 */

package _207_uniqueOccurrences

import (
	"fmt"
	"testing"
)

/**
给你一个整数数组 arr，请你帮忙统计数组中每个数的出现次数。

如果每个数的出现次数都是独一无二的，就返回 true；否则返回 false。
 */

func TestUniqueOccurrences(t *testing.T) {
	fmt.Println(uniqueOccurrences([]int{1,2}))
}

func uniqueOccurrences(arr []int) bool {

	// number of times statistics appear
	var numStastic map[int]int
	numStastic = make(map[int]int, len(arr))

	for _, v := range arr {
		numStastic[v]++
	}

	// number of times statistics appear for value
	var uniqueNum map[int]struct{}
	uniqueNum = make(map[int]struct{}, len(numStastic))

	for _, v := range numStastic {
		uniqueNum[v] = struct{}{}
	}

	return len(uniqueNum) == len(numStastic)
}
