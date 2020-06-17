package main

/**
 * @Auther: zhoujinxin@bytedance.com
 * @Date: 2020/6/15 22:13
 */

func main() {
	groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
}

func groupAnagrams(strs []string) [][]string {
	m := make(map[int32][]string)

	for _, str := range strs {
		var res int32
		for _, c := range str {
			res += c
		}

		if _, ok := m[res]; !ok {
			m[res] = []string{str}
			continue
		}

		tmp := m[res]
		tmp = append(tmp, str)
		m[res] = tmp
	}

	var result [][]string
	for _, value := range m {
		result = append(result, value)
	}
	return result
}
