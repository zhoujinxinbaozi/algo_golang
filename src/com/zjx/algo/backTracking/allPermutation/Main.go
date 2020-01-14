package main

import "fmt"

/***
	problem:

	给定一个没有重复数字的序列，返回其所有可能的全排列。

	示例:

	输入: [1,2,3]
	输出:
	[
	  [1,2,3],
	  [1,3,2],
	  [2,1,3],
	  [2,3,1],
	  [3,1,2],
	  [3,2,1]
	]

	solution:
	无关痛痒
*/
func main() {
	result := permute([]int{1, 2, 3})
	fmt.Println(result)
}

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	solution(nums, &result, len(nums), 0)
	return result
}

/*
	递归函数
 */
func solution(nums []int, result *[][]int, length int, current int) {
	if current == length {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*result = append(*result, tmp)
	}
	for i := current; i < length; i++ {
		swap(&nums, i, current)
		solution(nums, result, length, current+1)
		swap(&nums, i, current)
	}
}
/*
	交换两个切片相应位置的值
 */
func swap(nums *[]int, i int, current int) {
	tmp := (*nums)[i]
	(*nums)[i] = (*nums)[current]
	(*nums)[current] = tmp
}
