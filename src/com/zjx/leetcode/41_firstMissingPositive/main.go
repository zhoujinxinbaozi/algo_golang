package main

import "fmt"

/**
 * @Auther: zhoujinxin@bytedance.com
 * @Date: 2020/6/14 18:40
 */

/**
给你一个未排序的整数数组，请你找出其中没有出现的最小的正整数。



示例 1:

输入: [1,2,0]
输出: 3
示例 2:

输入: [3,4,-1,1]
输出: 2
示例 3:

输入: [7,8,9,11,12]
输出: 1

solution:
	所有的位置上 放上对应下标值+1， 超过数组长度和小于0的置0

*/
func main() {
	fmt.Println(firstMissingPositive([]int{1, 1}))
}

func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == i+1 { // correct position
			continue
		}
		if nums[i] <= 0 || nums[i] > len(nums) { // illegal position
			nums[i] = 0
			continue
		}
		if nums[i] == nums[nums[i]-1] { // the same  避免死循环
			// 后面的置为0
			if i > nums[i]-1 {
				nums[i] = 0
			} else {
				nums[nums[i]-1] = 0
			}
			continue
		}
		// Swap
		nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		i--
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] == i+1 {
			continue
		}
		return i + 1
	}

	return len(nums) + 1
}
