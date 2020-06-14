package main

/**
 * @Auther: zhoujinxin@bytedance.com
 * @Date: 2020/6/14 10:14
 */

/**
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

你可以假设数组中无重复元素。

示例 1:

输入: [1,3,5,6], 5
输出: 2
示例 2:

输入: [1,3,5,6], 2
输出: 1
示例 3:

输入: [1,3,5,6], 7
输出: 4
示例 4:

输入: [1,3,5,6], 0
输出: 0
*/
func main() {

}

func searchInsert(nums []int, target int) int {
	begin := 0
	end := len(nums) - 1
	if target > nums[len(nums)-1] {
		return len(nums)
	}
	if target < nums[0] {
		return 0
	}

	for begin <= end {
		middle := (begin + end) / 2
		if nums[middle] == target {
			return middle
		} else if nums[middle] > target {
			// 前一个小于
			if middle != 0 && target > nums[middle-1] {
				return middle
			}
			end = middle - 1
		} else {
			// 后一个大于
			if middle != len(nums)-1 && target < nums[middle+1] {
				return middle + 1
			}
			begin = middle + 1
		}
	}
	return -1
}
