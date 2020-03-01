package main

import "fmt"

/**
	problem:
	给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
	不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。

	solution:

*/
func main() {
	arr := removeDuplicates([]int{1, 1, 2})
	fmt.Println(arr)
}

func removeDuplicates(nums []int) int {
	len := len(nums)
	begin := 0
	for i := 1; i < len; i++ {
		if nums[i] != nums[begin] {
			begin++
			nums[begin] = nums[i]
		}
	}
	//fmt.Println(nums)
	return begin+1
}
