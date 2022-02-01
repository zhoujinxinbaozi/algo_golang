/**
 *  @Author: JinxinZhou
 *  @Date  : 2022/2/1 10:52 上午
 */

/*
搜索旋转数组。给定一个排序后的数组，包含n个整数，但这个数组已被旋转过很多次了，次数不详。
请编写代码找出数组中的某个元素，假设数组元素原先是按升序排列的。若有多个相同元素，返回索引值最小的一个。
旋转多次后可以保证二分之后的某一部分有序
*/

package _0_03_search

import (
	"fmt"
	"testing"
)

// TestSearch
func TestSearch(t *testing.T) {
	fmt.Println(search([]int{4, 5, 6, 7, 1, 2, 3}, 6))
}

func search(arr []int, target int) int {
	result := -1
	start := 0
	end := len(arr) - 1
	for start <= end {
		// 特殊处理 防止出现重复导致不准的情况
		// 5 1 2 3 5  没有当前这个判断会返回最后 一个5
		if arr[start] == target {
			return start
		}
		mid := (start + end) / 2
		// 等于结果 继续向前搜索是否还存在答案  因为要寻找最小的下标
		if arr[mid] == target {
			result = mid
			end = mid - 1
		} else if arr[mid] > arr[start] { // 左边有序
			if target >= arr[start] && target <= arr[mid] { // 结果在左边
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else if arr[mid] < arr[end] { // 右边有序
			if target >= arr[mid] && target <= arr[end] { // 结果在右边
				start = mid + 1
			} else {
				end = mid - 1
			}
		} else {
			start = start + 1
		}
	}
	return result
}
