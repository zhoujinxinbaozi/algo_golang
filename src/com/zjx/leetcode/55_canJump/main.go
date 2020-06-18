package main

import (
	"fmt"
	"sort"
)

/**
 * @Auther: zhoujinxin@bytedance.com
 * @Date: 2020/6/18 09:56
 */

/**
给定一个非负整数数组，你最初位于数组的第一个位置。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

判断你是否能够到达最后一个位置。

示例 1:

输入: [2,3,1,1,4]
输出: true
解释: 我们可以先跳 1 步，从位置 0 到达 位置 1, 然后再从位置 1 跳 3 步到达最后一个位置。
示例 2:

输入: [3,2,1,0,4]
输出: false
解释: 无论怎样，你总会到达索引为 3 的位置。但该位置的最大跳跃长度是 0 ， 所以你永远不可能到达最后一个位置。
*/
func main() {
	canJump([]int{2, 0, 0})
}

func canJump(nums []int) bool {
	resultArr := make([]int, len(nums))
	resultArr[0] = nums[0]
	if len(nums) == 1 {
		return true
	}
	for index, value := range nums { // 遍历每一个位置
		canReach := false
		for i := index - 1; i >= 0; i-- { // 查找之前的位置
			if index-i <= nums[i] && resultArr[i] != 0 { // 之前的某一个位置可以跳到当前位置，并且之前的某位置也可以到达
				canReach = true
				break
			}
		}

		if canReach {
			resultArr[index] = index + value
		}
	}
	fmt.Println(resultArr)
	sort.Ints(resultArr)
	fmt.Println(resultArr)
	return resultArr[len(nums)-1] >= len(nums)-1
}
