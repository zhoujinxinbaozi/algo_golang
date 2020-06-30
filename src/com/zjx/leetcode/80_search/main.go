package main

import "fmt"

/**
 * @Auther: zhoujinxin@bytedance.com
 * @Date: 2020/6/27 10:21
 */

/**
假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,0,1,2,2,5,6] 可能变为 [2,5,6,0,0,1,2] )。

编写一个函数来判断给定的目标值是否存在于数组中。若存在返回 true，否则返回 false。

示例 1:

输入: nums = [2,5,6,0,0,1,2], target = 0
输出: true
示例 2:

输入: nums = [2,5,6,0,0,1,2], target = 3
输出: false


同33题
*/
func main() {
	b := search([]int{2, 2, 2, 0, 2, 2}, 0)
	fmt.Println(b)
}

func search(nums []int, target int) bool {
	begin, end := 0, len(nums)-1

	for begin <= end {
		middle := (begin + end) / 2
		if nums[middle] == target {
			return true
		} else if (hasSort(nums, begin, middle) && target >= nums[begin] && target < nums[middle]) ||
			(hasSort(nums, middle, end) && !(target > nums[middle] && target <= nums[end])) {
			end = middle - 1
		} else {
			begin = middle + 1
		}
	}
	return false
}

func hasSort(nums []int, i, j int) bool {
	for i < j {
		if nums[i] <= nums[i+1] {
			i++
		} else {
			return false
		}
	}
	return true
}
