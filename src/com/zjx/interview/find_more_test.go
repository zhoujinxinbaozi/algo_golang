/**
 *  @Author: JinxinZhou
 *  @Date  : 2021/9/4 2:07 下午
 */

//
package interview

import (
	"fmt"
	"testing"
)

// 找到数组中第一个比给定数字大的数  不存在返回-1
func findMore(arr []int, target int) int {
	begin, end := 0, len(arr)-1
	result := -1
	for begin <= end {
		mid := (begin + end) / 2
		if arr[mid] <= target {
			begin = mid + 1
		} else {
			result = arr[mid]
			end = mid - 1
		}
	}
	return result
}

func TestFindMore(t *testing.T) {
	arr := []int{2, 3, 4, 6, 8, 9}
	target := 6
	fmt.Println(findMore(arr, target))
}
