package main

import "fmt"

// 最短无连续子数组
/**
problem:

给定一个整数数组，你需要寻找一个连续的子数组，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。

你找到的子数组应是最短的，请输出它的长度。

示例 1:

输入: [2, 6, 4, 8, 10, 9, 15]
输出: 5
解释: 你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。
说明 :

输入的数组长度范围在 [1, 10,000]。
输入的数组可能包含重复元素 ，所以升序的意思是<=。

solution:
首先找到发生逆序的左边界和右边界，这区间的数字是需要调整的,
在这区间找到最小值和最大值，最小值和前面的有序数列比，找到第一个大于min的数的位置
						 最大值和后面的有序数列比，找到最后一个比max大的数


*/
func main() {
	//findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 5})
	//findUnsortedSubarray([]int{1, 2, 3, 4, 5, 6, 7})
	findUnsortedSubarray([]int{9, 10, 1, 2, 3, 4})
}

func findUnsortedSubarray(nums []int) int {
	var i, left, right int
	// 发生逆序的左边界
	for i = 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			left = i
			break
		}
	}
	//发生逆序的右边界
	for i = len(nums) - 1; i > 0; i-- {
		if nums[i] < nums[i-1] {
			right = i
			break
		}
	}
	fmt.Printf("left=%d", left)
	fmt.Printf("right=%d", right)
	// 逆序区间中的min, max
	min, max := nums[left], nums[left]
	for i = left + 1; i <= right; i++ {
		if nums[i] < min {
			min = nums[i]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	// 逆序中的最大与后面的递增序列进行比较
	// 逆序中的最小与前面的递增序列进行比较
	var trueLeft, trueRight int
	for i = 0; i <= left; i++ {
		if min < nums[i] {
			trueLeft = i
			break
		}
	}
	for i = len(nums) - 1; i >= right; i-- {
		if max > nums[i] {
			trueRight = i
			break
		}
	}

	fmt.Printf("trueLeft=%d\n", trueLeft)
	fmt.Printf("trueRight=%d\n", trueRight)
	if trueLeft == trueRight {
		return 0
	} else {
		return trueRight - trueLeft + 1
	}
}
