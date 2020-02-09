package main

import (
	"fmt"
	"math"
)

/**
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。

示例 1:

nums1 = [1, 3]
nums2 = [2]

则中位数是 2.0
示例 2:

nums1 = [1, 2]
nums2 = [3, 4]

则中位数是 (2 + 3)/2 = 2.5

solution:
	https://blog.csdn.net/chen_xinjia/article/details/69258706
*/
func main() {
	fmt.Println(findMedianSortedArrays([]int{0, 1}, []int{3, 4, 5, 6}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	if len(nums2) < len(nums1) {
		return findMedianSortedArrays(nums2, nums1)
	}
	length := len(nums1) + len(nums2)
	cut1 := len(nums1) / 2  // nums1，切那一刀左边有几个元素
	cut2 := length/2 - cut1 // nums2，切那一刀左边有几个元素
	cutL := 0               // 在要切这一刀nums1的左边界
	cutR := len(nums1)      // 在要切这一刀nums1的右边界  方便用二分查找的方式加快速度
	for {
		if cut1 > len(nums1) {
			break
		}
		cut1 = (cutR-cutL)/2 + cutL
		cut2 = length/2 - cut1
		/**				L1   R1
				nums1 3  5 | 8  9
				nums2 1  2  7 | 10  11  12
		                    L2  R2
		正确切分后的四个数
		*/
		L1 := math.MinInt64
		L2 := math.MinInt64
		R1 := math.MaxInt64
		R2 := math.MaxInt64
		if cut1 != 0 {
			L1 = nums1[cut1-1]
		}
		if cut2 != 0 {
			L2 = nums2[cut2-1]
		}
		if cut1 != len(nums1) {
			R1 = nums1[cut1]
		}
		if cut2 != len(nums2) {
			R2 = nums2[cut2]
		}
		if L1 > R2 {
			cutR = cut1 - 1
		} else if L2 > R1 {
			cutL = cut1 + 1
		} else {
			if length%2 == 0 {
				if L1 < L2 {
					L1 = L2
				}
				if R1 > R2 {
					R1 = R2
				}
				return float64(L1+R1) / 2
			} else {
				if R1 < R2 {
					return float64(R1)
				} else {
					return float64(R2)
				}
			}
		}

	}
	return float64(-1)

}
