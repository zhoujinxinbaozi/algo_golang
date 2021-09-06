package _6_permute

/**
 * @Auther: zhoujinxin@bytedance.com
 * @Date: 2020/6/15 21:57
 */
/**
全排列
*/

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	solution(nums, &result, len(nums), 0)
	return result
}

func solution(nums []int, result *[][]int, length int, current int) {
	if current == length {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*result = append(*result, tmp)
	}
	for i := current; i < length; i++ {
		nums[i], nums[current] = nums[current], nums[i]
		solution(nums, result, length, current+1)
		nums[i], nums[current] = nums[current], nums[i]
	}
}
