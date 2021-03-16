/**
 *  @Author: JinxinZhou
 *  @Date  : 2021/3/15 7:45 下午
 */

package recursion

import (
	"fmt"
	"testing"
)

// TestFindMaxDivide 给你一块 X * Y 的地，均匀的分成正方块，分出的方块要尽量大
func TestFindMaxDivide(t *testing.T) {
	x, y := 168, 15
	fmt.Println(Divide(x, y))
}

func Divide(x, y int) int {
	if x > y {
		Divide(y, x)
	}

	if y%x == 0 {
		return x
	}

	return Divide(y-(x*(y/x)), x)
}

// TestPaiLie 数组的全排列
func TestPaiLie(t *testing.T) {
	arr := []int{1, 2, 3}
	QuanPaiLie(arr, 0)
}

func QuanPaiLie(arr []int, cur int) {
	if cur == (len(arr) - 1) {
		fmt.Println(arr)
		return
	}
	for i := cur; i < len(arr); i++ {
		arr[i], arr[cur] = arr[cur], arr[i]
		QuanPaiLie(arr, cur+1)
		arr[i], arr[cur] = arr[cur], arr[i]
	}
}
