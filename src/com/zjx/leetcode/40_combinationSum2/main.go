package main

import (
	"fmt"
	"sort"
)

/**
 * @Auther: zhoujinxin@bytedance.com
 * @Date: 2020/6/14 16:12
 */

/**
给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用一次。

说明：

所有数字（包括目标数）都是正整数。
解集不能包含重复的组合。
示例 1:

输入: candidates = [10,1,2,7,6,1,5], target = 8,
所求解集为:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
示例 2:

输入: candidates = [2,5,2,1,2], target = 5,
所求解集为:
[
  [1,2,2],
  [5]
]
*/
func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

func combinationSum2(candidates []int, target int) [][]int {
	var result [][]int
	fmt.Println(candidates)
	sort.Ints(candidates)
	fmt.Println(candidates)
	RecusiveCombinationSum(candidates, 0, 0, target, []int{}, &result)
	return result
}

func RecusiveCombinationSum(candidates []int, begin, sum int, target int, list []int, result *[][]int) {
	if sum == target {
		list1 := make([]int, len(list))
		copy(list1, list)
		if !ContainsSlice(*result, list1) {
			*result = append(*result, list1)
		}
		return
	}
	if sum > target {
		return
	}

	for i := begin; i < len(candidates); i++ {
		list = append(list, candidates[i])
		RecusiveCombinationSum(candidates, i+1, sum+candidates[i], target, list, result)
		list = list[:len(list)-1]
	}
}

// result contain target
func ContainsSlice(result [][]int, target []int) bool {

	for _, slice := range result {
		var has = true
		for index, v := range slice {
			if len(slice) != len(target) {
				has = false
				break
			}
			if v != target[index] {
				has = false
			}
		}
		if has {
			return true
		}
	}
	return false
}
