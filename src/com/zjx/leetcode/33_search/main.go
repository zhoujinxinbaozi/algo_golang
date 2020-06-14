package main

import "fmt"

/**
 * @Auther: zhoujinxin@bytedance.com
 * @Date: 2020/6/13 18:15
 */

/*
假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。

你可以假设数组中不存在重复的元素。

你的算法时间复杂度必须是 O(log n) 级别。
*/
func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}

func search(nums []int, target int) int {
	return BinarySearch(nums, 0, len(nums)-1, target)
}

func BinarySearch(nums []int, begin, end, target int) int {

	for begin <= end {
		middle := (begin + end) / 2

		if nums[middle] == target {
			return middle
			// 左半部分有序，并在其中
			// 右半部分有序，并且不在其中
		} else if nums[begin] <= nums[middle] && target >= nums[begin] && target < nums[middle] ||
			(nums[middle] <= nums[end] && !(target > nums[middle] && target <= nums[end])) {

			end = middle - 1
		} else {
			begin = middle + 1
		}
	}
	return -1
}
