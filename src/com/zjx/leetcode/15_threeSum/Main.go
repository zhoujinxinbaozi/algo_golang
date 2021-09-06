package main

import (
	"fmt"
	"sort"
)

/**
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

 

示例：

给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]

solution:
	对数组进行排序
	当前位置为i， 下一个元素为j，数组的末尾元素为k
	if nums[i] + nums[j] + nums[k] > 0  k --
	else if                        < 0  j ++
	else                                加入到结果集中 k-- j++
	去重

*/
func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		j := i + 1
		k := len(nums) - 1
		for {
			if j >= k {
				break
			}
			if nums[i]+nums[j]+nums[k] > 0 {
				k--
			} else if nums[i]+nums[j]+nums[k] < 0 {
				j++
			} else {
				target := []int{nums[i], nums[j], nums[k]}
				sort.Ints(target)
				add(&result, target)
				k--
				j++
			}
		}

	}
	return result
}

// 判断result中是否含有target，没有则加入到result
func add(result *[][]int, target []int) {
	a := false
	for _, value := range *result {
		b := true
		for i := 0; i < 3; i++ {
			if target[i] != value[i] {
				b = false
				break
			}
		}
		if b {
			a = true
			break
		}
	}
	if !a {
		*result = append(*result, target)
	}
}
