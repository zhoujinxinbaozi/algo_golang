package main

import "fmt"

/**
 * @Auther: zhoujinxin@bytedance.com
 * @Date: 2020/6/17 19:10
 */

/**
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
*/
func main() {
	fmt.Println(maxSubArray1([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

// 使用reslut数组存储以当前位置i结尾所获得的最大值，result[i] = max(nums[i], nums[i]+result[i-1]) // 不适用前面的结果，使用前面的结果
func maxSubArray(nums []int) int {
	result := make([]int, len(nums))
	result[0] = nums[0]
	singleResult := nums[0]

	for i := 1; i < len(nums); i++ {
		result[i] = max(nums[i], nums[i]+result[i-1])
		singleResult = max(result[i], singleResult)
	}
	fmt.Println(result)
	return singleResult
}

// 使用bmax存放当前位置之前的累加最大和
func maxSubArray1(nums []int) int {
	bmax := nums[0]
	singleResult := nums[0]

	for i := 1; i < len(nums); i++ {
		bmax = max(bmax+nums[i], nums[i])
		singleResult = max(singleResult, bmax)
	}
	return singleResult
}

func max(i int, i2 int) int {
	if i < i2 {
		return i2
	}
	return i
}
