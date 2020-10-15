/**
 *  @Author: JinxinZhou
 *  @Date  : 2020/10/14 7:22 下午
 */

package _91_findSubsequences

import (
	"fmt"
	"testing"
)

var result [][]int

func findSubsequences(nums []int) [][]int {
	for i := 0; i < len(nums); i++ {
		dfs(i, nums, []int{})
	}
	return result
}

func dfs(i int, nums, tmp []int) {

	if len(tmp) >= 2 && !judgeContain(tmp) {
		fmt.Println("before : ", result)
		var r []int
		r = make([]int, len(tmp))
		copy(r, tmp)
		result = append(result, r)
		fmt.Println("after : ", result)
	}

	for ; i < len(nums); i++ {
		var add bool
		if len(tmp) == 0 {
			tmp = append(tmp, nums[i])
		} else {
			if nums[i] >= tmp[len(tmp)-1] {
				tmp = append(tmp, nums[i])
				add = true
			}
		}
		dfs(i+1, nums, tmp)
		if add {
			tmp = tmp[0:len(tmp)-1]
		}
	}

}



func judgeContain(tmp []int) bool {
	var r bool

	for _, arr := range result {
		if len(arr) != len(tmp) {
			continue
		}

		var count int
		for i := 0; i < len(tmp); i++ {
			if arr[i] == tmp[i] {
				count++
			}
		}

		if count == len(tmp) {
			r = true
			break
		}

	}
	return r
}

func TestFindSubsequences(t *testing.T) {
	fmt.Print("lll ", findSubsequences([]int{1, 3, 5, 7}))
}
