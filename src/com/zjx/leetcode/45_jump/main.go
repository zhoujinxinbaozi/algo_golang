package main

import (
	"math"
)

/**
 * @Auther: zhoujinxin@bytedance.com
 * @Date: 2020/6/15 21:30
 */

/**
给定一个非负整数数组，你最初位于数组的第一个位置。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

你的目标是使用最少的跳跃次数到达数组的最后一个位置。

示例:

输入: [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。

*/

func main() {
	jump([]int{2, 3, 1, 1, 4})
}

func jump(nums []int) int {
	arr := make([]int, len(nums)) // 到达当前位置的最小跳数
	arr[0] = 0
	for i := 1; i < len(nums); i++ {
		arr[i] = math.MaxInt64
	}

	for i := 1; i < len(nums); i++ { // i n-1
		for j := i - 1; j >= 0; j-- { // i 之前所有
			if j+nums[j] >= i {
				arr[i] = min(arr[i], arr[j]+1)
			}
		}
		// arr[i] 是单调性的，上面的for可以剪枝
		//for j := 0; j < i; j++ {
		//	if j+nums[j] >= i {
		//		arr[i] = min(arr[i], arr[j]+1)
		//		break
		//	}
		//}

	}
	return arr[len(arr)-1]
}

func min(i int, j int) int {
	if i < j {
		return i
	}
	return j
}
