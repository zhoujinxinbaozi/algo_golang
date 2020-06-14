package main

/**
 * @Auther: zhoujinxin@bytedance.com
 * @Date: 2020/6/13 19:21
 */

/**
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

你的算法时间复杂度必须是 O(log n) 级别。

如果数组中不存在目标值，返回 [-1, -1]。

示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]
示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]
*/

func main() {
	searchRange([]int{5, 7, 7, 8, 8, 10}, 8)
}

func searchRange(nums []int, target int) []int {

	return SearchIndex(nums, 0, len(nums)-1, target)
}

// 找到存在的元素的下标
func SearchIndex(nums []int, begin, end, target int) []int {
	location := -1
	for begin <= end {
		middle := (begin + end) / 2
		if nums[middle] == target {
			location = middle
			break
		} else if nums[middle] > target {
			end = middle - 1
		} else {
			begin = middle + 1
		}
	}

	if location == -1 {
		return []int{-1, -1}
	}

	return FindFirstAndLast(nums, location)
}

// 找到前后的index
func FindFirstAndLast(nums []int, location int) []int {
	leftIndex, rightIndex := location, location
	for i := location + 1; i < len(nums); i++ {
		if nums[i] != nums[location] {
			break
		}
		rightIndex = i
	}

	for i := location - 1; i >= 0; i-- {
		if nums[i] != nums[location] {
			break
		}
		leftIndex = i
	}

	return []int{leftIndex, rightIndex}
}
