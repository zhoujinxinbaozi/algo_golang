/**
 *  @Author: JinxinZhou
 *  @Date  : 2020/10/25 7:04 下午
 */

package _96_nextGreaterElement

import (
	"fmt"
	"testing"
)

func TestNextGreaterElement(t *testing.T) {
	result := nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2})
	fmt.Println(result)
}

//func nextGreaterElement(nums1 []int, nums2 []int) []int {
//	var result []int
//	for _, n1 := range nums1 {
//		var isReach bool
//		var isThrough bool
//		for _, n2 := range nums2 {
//			if n1 != n2 && !isReach{
//				continue
//			}
//			isReach = true
//			if n2 > n1 {
//				result = append(result, n2)
//				isThrough = true
//				break
//			}
//		}
//
//		if !isThrough {
//			result = append(result, -1)
//		}
//
//	}
//	return result
//}

// 单调栈解法
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var result []int
	var tmp []int
	var m map[int]int
	m = make(map[int]int)

	for i := len(nums2) - 1; i >= 0; i-- {

		if len(tmp) == 0 {
			//m[nums2[i]] = math.MaxInt32
			tmp =append(tmp, nums2[i])
			continue
		}
		if nums2[i] < tmp[len(tmp)-1] {
			tmp = append(tmp, nums2[i])
			m[nums2[i]] = tmp[len(tmp)-2]
			continue
		}

		for len(tmp) != 0 && nums2[i] > tmp[len(tmp)-1] {
			tmp = tmp[:len(tmp)-1]
		}
		tmp = append(tmp, nums2[i])
		if len(tmp) == 1 {
			continue
		}
		m[nums2[i]] = tmp[len(tmp)-2]
	}

	for _, n1 := range nums1 {
		if v, ok := m[n1]; ok {
			result = append(result, v)
			continue
		}
		result = append(result, -1)
	}

	return result
}
