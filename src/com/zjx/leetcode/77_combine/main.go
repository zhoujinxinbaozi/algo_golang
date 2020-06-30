package main

/**
 * @Auther: zhoujinxin@bytedance.com
 * @Date: 2020/6/27 09:11
 */

/**
给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

示例:

输入: n = 4, k = 2
输出:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
*/
func main() {
	combine(4, 2)
}

func combine(n int, k int) [][]int {

	var result [][]int

	GetAnswer(n, 1, 0, k, []int{}, &result)
	//fmt.Println(result)
	return result
}

func GetAnswer(n int, begin, current int, k int, arr []int, result *[][]int) {
	if current == k {
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		*result = append(*result, tmp)
		return
	}

	for i := begin; i <= n; i++ {
		arr = append(arr, i)
		//fmt.Println(arr)
		GetAnswer(n, i+1, current+1, k, arr, result)
		arr = arr[:len(arr)-1]
		//fmt.Println(arr)
	}
}
