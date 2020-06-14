package main

import "fmt"

/**
 * @Auther: zhoujinxin@bytedance.com
 * @Date: 2020/6/14 10:34
 */

/**
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：

所有数字（包括 target）都是正整数。
解集不能包含重复的组合。
示例 1:

输入: candidates = [2,3,6,7], target = 7,
所求解集为:
[
  [7],
  [2,2,3]
]
示例 2:

输入: candidates = [2,3,5], target = 8,
所求解集为:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]
*/
func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	RecusiveCombinationSum(candidates, 0, 0, target, []int{}, &result)
	return result
}

func RecusiveCombinationSum(candidates []int, begin, sum int, target int, list []int, result *[][]int) {
	if sum == target {
		list1 := make([]int, len(list))
		copy(list1, list)
		*result = append(*result, list1)
		return
	}
	if sum > target {
		return
	}

	for i := begin; i < len(candidates); i++ {
		list = append(list, candidates[i])
		RecusiveCombinationSum(candidates, i, sum+candidates[i], target, list, result)
		list = list[:len(list)-1]
	}
}
