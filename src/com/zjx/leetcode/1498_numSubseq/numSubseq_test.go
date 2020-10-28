/**
 *  @Author: JinxinZhou
 *  @Date  : 2020/10/27 2:55 下午
 */

package _498_numSubseq

import (
	"fmt"
	"testing"
)

var result int
var tmp []int

func TestNumSubSeq(t *testing.T) {
	fmt.Println(numSubseq([]int{3, 5, 6, 7}, 9))
}

func numSubseq(nums []int, target int) int {
	result = 0
	tmp = nil
	bfs(nums, target, 0)
	return result
}

func bfs(nums []int, target int, index int) {

	getCurrentResult(target)

	for ; index < len(nums); index++ {
		tmp = append(tmp, nums[index])
		bfs(nums, target, index+1)
		tmp = tmp[:len(tmp)-1]
	}
}

func getCurrentResult(target int) {
	if tmp == nil {
		return
	}

	var min, max int
	min, max = tmp[0], tmp[0]
	for _, v := range tmp {
		if v < min {
			min = v
		}

		if v > max {
			max = v
		}
	}

	if min+max <= target {
		result++
	}
}
