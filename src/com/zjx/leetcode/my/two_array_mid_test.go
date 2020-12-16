/**
 *  @Author: JinxinZhou
 *  @Date  : 2020/12/15 5:11 下午
 */

package my

import (
	"fmt"
	"testing"
)

func TestGetKthElement(t *testing.T) {
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var result float64
	length := len(nums2) + len(nums1)
	if length&1 == 1 {
		result = float64(KElement(0, 0, nums1, nums2, length/2+1))
	} else {
		result = (float64(KElement(0, 0, nums1, nums2, length/2)) +
			float64(KElement(0, 0, nums1, nums2, length/2+1))) / 2
	}
	return result
}

// KElement 获取两个有序数组中第k小的数
// @param begin1  nums1起始位置
// @param begin2  nums2起始位置
// @param nums1,nums2 列表
// @param k       第k小的数
// @return int    第k小的数值
func KElement(begin1, begin2 int, nums1, nums2 []int, k int) int {
	if len(nums1) == 0 {
		return nums2[k-1]
	}

	if len(nums2) == 0 {
		return nums1[k-1]
	}

	if begin1 == len(nums1) {
		return nums2[begin2+k-1]
	}

	if begin2 == len(nums2) {
		return nums1[begin1+k-1]
	}

	if k == 1 {
		return min(nums1[begin1], nums2[begin2])
	}

	nums1Index := max(min(begin1+k/2, len(nums1))-1, 0)
	nums2Index := max(min(begin2+k/2, len(nums2))-1, 0)

	nums1Val, nums2Val := nums1[nums1Index], nums2[nums2Index]
	if nums1Val <= nums2Val {
		return KElement(nums1Index+1, begin2, nums1, nums2, k-(nums1Index-begin1+1))
	} else {
		return KElement(begin1, nums2Index+1, nums1, nums2, k-(nums2Index-begin2+1))
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
