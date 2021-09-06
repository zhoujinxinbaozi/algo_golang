/**
 *  @Author: JinxinZhou
 *  @Date  : 2021/12/24 10:09 下午
 */

package interview

import (
	"fmt"
	"math"
	"testing"
)

var maxLength int

func TestK(t *testing.T) {
	maxLength = math.MinInt32
	arr := []int{1, -1, 5,-2,3}
	getResult(arr, 0, len(arr), 3, 0, 0)
	fmt.Println(maxLength)
}

func getResult(arr []int, start, end, k, count, countSum int) {
	if countSum == k {
		maxLength = max(maxLength, count)
	}
	for i := start; i < end; i++ {
		if i >= len(arr) {
			continue
		}
		getResult(arr, i+1, i+3, k, count+1, countSum + arr[i])
	}
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func TestPre(t *testing.T) {
	//arr := []int{1, -1, 5, -2, 3}
	arr := []int{-2, -1, 2, 1}
	m := make(map[int]int)
	//k := 3
	k := 1
	count := 0
	max := math.MinInt32
	for index, num := range arr {
		if _, ok := m[k-num]; !ok {
			count += num
			if _, ok := m[count]; !ok {
				m[count] = index + 1
			}
			m[k-num] = index + 1
		}
		if _, ok := m[num]; !ok {
			m[num] = index + 1
		}
		//if _, ok := m[k-num]; !ok {
		//	count += num
		//	m[count] = index+1
		//	if _, ok := m[num]; !ok {
		//		m[num] = index + 1
		//	}
		//	continue
		//}
		//if max < num - m[k-num] {
		//	fmt.Printf("index : %v, m : %v\n", index, m[k-num])
		//	max = index - m[k-num]
		//}
		//if _, ok := m[num]; !ok {
		//	m[num] = index + 1
		//}
	}
	fmt.Printf("mm : %v\n", m)
	fmt.Printf("max : %v\n", max+1)
}
