/**
 *  @Author: JinxinZhou
 *  @Date  : 2020/10/28 2:34 下午
 */

/**
给定两个字符串 s1 和 s2，请编写一个程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。
 */
package _102_CheckPermutation

func CheckPermutation(s1 string, s2 string) bool {
	var m map[int32]int
	m = make(map[int32]int, len(s1))

	for _, v := range s1 {
		if _, ok := m[v]; ok {
			m[v]++
			continue
		}
		m[v] = 1
	}

	for _, v := range s2 {
		if _, ok := m[v]; !ok {
			return false
		}
		if m[v] <= 0 {
			return false
		}
		m[v]--
	}

	return true
}
