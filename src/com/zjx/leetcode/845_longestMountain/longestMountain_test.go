/**
 *  @Author: JinxinZhou
 *  @Date  : 2020/10/25 2:15 下午
 */

package _45_longestMountain

import (
	"fmt"
	"testing"
)

func TestLongestMountain(t *testing.T) {
	result := longestMountain([]int{2, 1, 4, 7, 3, 2, 5})
	fmt.Println(result)
}

func longestMountain(A []int) int {
	if A == nil {
		return 0
	}
	var result int
	var leftExtend []int
	for i := 0; i < len(A); i++ {
		if i == 0 {
			leftExtend = append(leftExtend, 0)
			continue
		}
		if A[i] > A[i-1] {
			if leftExtend[len(leftExtend)-1] == 0 {
				leftExtend = append(leftExtend, leftExtend[len(leftExtend)-1]+2)
			} else {
				leftExtend = append(leftExtend, leftExtend[len(leftExtend)-1]+1)
			}
		} else {
			leftExtend = append(leftExtend, 0)
		}
	}

	var rightExtend []int
	for i := len(A) - 1; i >= 0; i-- {
		if i == len(A)-1 {
			rightExtend = append(rightExtend, 0)
			continue
		}
		if A[i] > A[i+1] {
			if rightExtend[len(rightExtend)-1] == 0 {
				rightExtend = append(rightExtend, rightExtend[len(rightExtend)-1]+2)
			} else {
				rightExtend = append(rightExtend, rightExtend[len(rightExtend)-1]+1)
			}
		} else {
			rightExtend = append(rightExtend, 0)
		}
	}
	intReverse(rightExtend)

	for i := 0; i < len(A); i++ {
		result = max(result, leftExtend[i]+rightExtend[i]-1)
	}

	return result
}

func intReverse(src []int) {
	if src == nil {
		panic(fmt.Errorf("the src can't be empty!"))
	}
	count := len(src)
	mid := count / 2
	for i := 0; i < mid; i++ {
		tmp := src[i]
		src[i] = src[count-1]
		src[count-1] = tmp
		count--
	}
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
