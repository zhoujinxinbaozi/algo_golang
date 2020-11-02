/**
 *  @Author: JinxinZhou
 *  @Date  : 2020/11/2 4:48 下午
 */

package _49_intersection

func intersection(nums1 []int, nums2 []int) []int {
	l1 := len(nums1)
	l2 := len(nums2)
	if l1 > l2 {
		return cal(nums1, nums2)
	} else {
		return cal(nums2, nums1)
	}
}

func cal(nums1 []int, nums2 []int) []int {
	m := map[int]struct{}{}
	// 去重map
	contain := map[int]struct{}{}
	for _, v := range nums1 {
		m[v] = struct{}{}
	}
	var result []int
	for _, v := range nums2 {
		if _, ok := m[v]; ok {
			if _, o := contain[v]; !o{
				result = append(result, v)
				contain[v] = struct{}{}
			}
		}
	}
	return result
}
