/**
 *  @Author: JinxinZhou
 *  @Date  : 2021/12/1 2:46 下午
 */

package merge_two_list

import (
	"fmt"
	"math"
	"testing"
)

type Node struct {
	Val int
	Next *Node
}


func TestHh(t *testing.T) {
	list1 := &Node{Val: 1}
	listNode1 := &Node{Val: 2}
	listNode2 := &Node{Val: 3}
	list1.Next = listNode1
	listNode1.Next = listNode2

	list2 := &Node{Val: 1}
	listNode3 := &Node{Val: 4}
	listNode4 := &Node{Val: 7}
	list2.Next = listNode3
	listNode3.Next = listNode4

	result := merge(list1, list2)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}

func merge(list1 *Node, list2 *Node) *Node {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var result *Node
	if list1.Val < list2.Val {
		result = list1
		result.Next = merge(list1.Next, list2)
	} else {
		result = list2
		result.Next = merge(list1, list2.Next)
	}
	return result
}
// 股票1   只允许买卖股票一次
func maxProfit(prices []int) int {
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}
	var result int
	min := math.MaxInt32 // 存贮到位置i的最小值
	for i := 0; i < len(prices); i++ {
		if min > prices[i] {
			min = prices[i]
		}
		if prices[i] - min > result {
			result = prices[i] - min
		}
	}
	return result
}

// 可以买卖多次的情况
func maxProfit2(prices []int) int {
	length := len(prices)
	var result int
	for i := 1; i < length; i++ {
		if prices[i] - prices[i-1] > 0 {
			result += (prices[i] - prices[i-1])
		}
	}
	return result
}
//func maxProfit(prices []int) int {
//	if len(prices) == 0 || len(prices) == 1 {
//		return 0
//	}
//	math.MaxInt32
//	var minSlice []int
//	var maxSlice []int
//	length := len(prices)
//	tmp := prices[0]
//	for i := 0; i < length; i++{
//		if prices[i] <= tmp {
//			minSlice = append(minSlice, prices[i])
//			tmp = prices[i]
//		} else {
//			minSlice = append(minSlice, tmp)
//		}
//	}
//	tmp = prices[length-1]
//	for i := length-1; i >= 0; i-- {
//		if prices[i] >= tmp {
//			maxSlice = append([]int{prices[i]}, maxSlice...)
//			tmp = prices[i]
//		} else {
//			maxSlice = append([]int{tmp}, maxSlice...)
//		}
//	}
//	var result int
//	for i := 0; i < length-1; i++ {
//		if maxSlice[i+1] - minSlice[i] > result {
//			result = maxSlice[i+1] - minSlice[i]
//		}
//	}
//	return result
//}