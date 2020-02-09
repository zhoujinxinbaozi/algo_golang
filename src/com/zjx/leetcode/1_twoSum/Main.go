package main

import "fmt"

/**
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

solution:
	用map存储，key为数组的值，value为数组的下表
	遍历数组时，当前值为current，就在map中查找时候有target-current，有则返回
*/
func main() {
	fmt.Println(twoSum([]int{3, 2, 3}, 6))
}

func twoSum(nums []int, target int) []int {
	var m map[int]int // 存储  key=nums[i] value=i
	m = make(map[int]int, len(nums))
	for index, value := range nums {
		if value, ok := m[target-value]; ok {
			return []int{index, value}
		}
		m[value] = index
	}
	return []int{0, 1}
}
