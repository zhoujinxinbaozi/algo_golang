/**
 *  @Author: JinxinZhou
 *  @Date  : 2020/10/25 2:15 下午
 */

package _45_longestMountain
/*
把符合下列属性的数组 arr 称为 山脉数组 ：

arr.length >= 3
存在下标 i（0 < i < arr.length - 1），满足
arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
arr[i] > arr[i + 1] > ... > arr[arr.length - 1]
给出一个整数数组 arr，返回最长山脉子数组的长度。如果不存在山脉子数组，返回 0 。

 

示例 1：

输入：arr = [2,1,4,7,3,2,5]
输出：5
解释：最长的山脉子数组是 [1,4,7,3,2]，长度为 5。
示例 2：

输入：arr = [2,2,2]
输出：0
解释：不存在山脉子数组。
 */
import (
	"fmt"
	"testing"
)

func TestLongestMountain(t *testing.T) {
	result := longestMountain([]int{2, 1, 4, 7, 3, 2, 5})
	fmt.Println(result)
}

func longestMountain(arr []int) int {
	left := make([]int,len(arr),len(arr)) // 当前位置左侧满足条件的最大个数
	for i := 1; i < len(arr); i++ {
		if arr[i] <= arr[i-1]{
			continue
		}
		left[i] = left[i-1] + 1
	}
	right := make([]int,len(arr),len(arr)) // 当前位置右侧满足条件的最大个数
	for i := len(arr)-2; i >= 0; i-- {
		if arr[i] <= arr[i+1]{
			continue
		}
		right[i] = right[i+1] + 1
	}
	result := 0
	for i := 1; i < len(arr); i++ {
		if left[i] == 0 || right[i] == 0 {
			continue
		}
		result = max(result,left[i]+right[i]+1)
	}
	if result >= 3 {
		return result
	}
	return 0

}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}