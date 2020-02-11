package main

import (
	"fmt"
	"sort"
)
/**
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

注意：

答案中不可以包含重复的四元组。

示例：

给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]

solution:
	同三个数，这次固定的是两个数即可
 */
func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}

func fourSum(nums []int, target int) [][]int {
	var result [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		j := i + 1
		for{
			if j >= len(nums)-2 {
				break
			}
			left := j + 1
			right := len(nums) - 1
			for {
				if left >= right {
					break
				}
				if nums[i]+nums[j]+nums[left]+nums[right] > target {
					right--
				} else if nums[i]+nums[j]+nums[left]+nums[right] < target {
					left++
				} else {
					target := []int{nums[i], nums[j], nums[left], nums[right]}
					sort.Ints(target)
					add(&result, target)
					left++
					right--
				}
			}
			j ++
		}
	}
	return result
}

// 判断result中是否含有target，没有则加入到result
func add(result *[][]int, target []int) {
	a := false
	for _, value := range *result {
		b := true
		for i := 0; i < len(target); i++ {
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
