/**
 *  @Author: JinxinZhou
 *  @Date  : 2021/5/31 3:51 下午
 */

package interview

import (
	"fmt"
	"testing"
)

func partition(arr []int, low, high int) int {
	pivot := arr[low] //导致 low 位置值为空
	for low < high {
		//high指针值 >= pivot high指针👈移
		for low < high && pivot <= arr[high] {
			high--
		}
		//填补low位置空值
		//high指针值 < pivot high值 移到low位置
		//high 位置值空
		arr[low] = arr[high]
		//low指针值 <= pivot low指针👉移
		for low < high && pivot >= arr[low] {
			low++
		}
		//填补high位置空值
		//low指针值 > pivot low值 移到high位置
		//low位置值空
		arr[high] = arr[low]
	}
	//pivot 填补 low位置的空值
	arr[low] = pivot
	fmt.Println(arr)
	return low
}

func QuickSort(arr []int, low, high int) {
	if low >= high {
		return
	}
	//位置划分
	pivot := partition(arr, low, high)
	//左边部分排序
	QuickSort(arr, low, pivot-1)
	//右边排序
	QuickSort(arr, pivot+1, high)
}

func TestQuickSort(t *testing.T) {
	arr := []int{3, 5, 4, 6, 2, 7}
	QuickSort(arr, 0, len(arr)-1)
	t.Log(arr)
}
