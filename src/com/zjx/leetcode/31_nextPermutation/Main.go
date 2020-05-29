package main

import (
	"sort"
)

/**
 * @Auther: zhoujinxin@bytedance.com
 * @Date: 2020/5/23 16:55
 */

func main() {
	nextPermutation([]int{1, 5, 1})
}
func nextPermutation(nums []int) {
	// 第一步从右往左查找最大的升序队列
	i := len(nums) - 2 // 倒数第二个数
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i == -1 {
		// 从右往左一直是升序
		sort.Ints(nums)
	} else {
		j := len(nums) - 1
		// 从右往左找第一个比i索引值大的数字索引坐标
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		//交换他两位置
		nums[i], nums[j] = nums[j], nums[i]
		// 重新排序
		sort.Ints(nums[i+1:])

	}
}
