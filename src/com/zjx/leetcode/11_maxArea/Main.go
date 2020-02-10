package main

import "fmt"
/**
给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。
示例:

输入: [1,8,6,2,5,4,8,3,7]
输出: 49

solution:
	两个指针，分别指向头left和尾right，取min(height[left], height[right]), min的值与right-left想乘。与全局的result的进行对比。
	if height[left] < height[right] left 右移，直到找到大于当前值得下标停止
	else                            right 左移，直到找到大于当前值得下标停止

 */
func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	result := 0
	left := 0
	right := len(height) - 1
	leftHight := 0
	rightHight := 0
	for {
		if left >= right {
			break
		}
		var minHight int
		leftHight = height[left]
		rightHight = height[right]
		if leftHight < rightHight {
			minHight = leftHight
		} else {
			minHight = rightHight
		}
		if (right-left)*minHight > result {
			result = (right - left) * minHight
		}
		if leftHight < rightHight {
			for {
				if left >= right {
					break
				}
				left++
				if height[left] > leftHight {
					break
				}
			}
		} else {
			for {
				if left >= right {
					break
				}
				right--
				if height[right] > rightHight {
					break
				}
			}
		}

	}
	return result
}
