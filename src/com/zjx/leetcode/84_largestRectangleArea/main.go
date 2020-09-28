/**
 *  @Author: JinxinZhou
 *  @Date  : 2020/9/28 10:59 上午
 */
/*
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
求在该柱状图中，能够勾勒出来的矩形的最大面积。
*/
/*
单调递增栈：
			可以找到左右第一个比他小的元素  栈内保持从栈底到栈顶为递增，否则出栈，当前的下一个元素为第一个比当前数小的
单调递减栈：
            可以找到左右第一个比他大的元素
 */
package main

import (
	"fmt"
)

func main() {
	result := largestRectangleArea([]int{1})
	fmt.Println(result)
}
type Node struct {
	Index int
	Value int
}
// increase 单调递增栈  当前位置左/右 第一个小于当前的值
func increaseLeft(arr []int) []int {
	var result []int

	var nodeList []Node
	for i, v := range arr {

		if !canPop(nodeList) {
			result = append(result, -1)
			nodeList = append(nodeList, Node{
				Index: i,
				Value: v,
			})
			continue
		}

		isOver := false
		for canPop(nodeList) {
			node := nodeList[len(nodeList)-1]
			if v > node.Value {
				nodeList = append(nodeList, Node{
					Index: i,
					Value: v,
				})
				result = append(result, node.Index)
				isOver = true
				break
			}else {
				nodeList = nodeList[0:len(nodeList)-1]
			}
		}
		if isOver {
			continue
		}
		result = append(result, -1)
		nodeList = append(nodeList, Node{
			Index: i,
			Value: v,
		})
	}
	return result
}

//slice翻转
func SliceIntReverse(src []int){
	if src == nil {
		return
	}
	count := len(src)
	mid := count/2
	for i := 0;i < mid; i++{
		tmp := src[i]
		src[i] = src[count-1]
		src[count-1] = tmp
		count--
	}
}

// increase 单调递增栈  当前位置左/右 第一个小于当前的值
func increaseRight(arr []int) []int {
	var result []int

	var nodeList []Node
	for i:=len(arr)-1; i >= 0; i -- {
		v := arr[i]
		if !canPop(nodeList) {
			result = append(result, len(arr))
			nodeList = append(nodeList, Node{
				Index: i,
				Value: v,
			})
			continue
		}

		isOver := false
		for canPop(nodeList) {
			node := nodeList[len(nodeList)-1]
			if v > node.Value {
				nodeList = append(nodeList, Node{
					Index: i,
					Value: v,
				})
				result = append(result, node.Index)
				isOver = true
				break
			}else {
				nodeList = nodeList[0:len(nodeList)-1]
			}
		}
		if isOver {
			continue
		}
		result = append(result, len(arr))
		nodeList = append(nodeList, Node{
			Index: i,
			Value: v,
		})
	}
	SliceIntReverse(result)
	return result
}

// canPop 是否可以弹出
func canPop(nodeList []Node) bool {
	return len(nodeList) > 0
}

// largestRectangleArea 以数组中的每一个为准，向左右分别找>=当前值得最远index，求面积取最大
func largestRectangleArea(arr []int) int {
	var result int
	increaseR := increaseRight(arr)
	increaseL := increaseLeft(arr)

	for i := 0; i < len(arr); i++ {

		// find left first less arr[i] index
		left := increaseL[i]

		// find right first more arr[i] index
		right := increaseR[i]

		tmp := (right-left-1)*arr[i]
		result = GetMax(result, tmp)
	}

	return result
}

// GetMax 获取最大值
func GetMax(left, right int) int {
	var result int
	if left < right {
		result = right
	} else {
		result = left
	}
	return result
}
