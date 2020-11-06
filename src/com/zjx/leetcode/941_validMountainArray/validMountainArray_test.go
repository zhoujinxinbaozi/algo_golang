/**
 *  @Author: JinxinZhou
 *  @Date  : 2020/11/3 7:43 下午
 */
/*
给定一个整数数组 A，如果它是有效的山脉数组就返回 true，否则返回 false。

让我们回顾一下，如果 A 满足下述条件，那么它是一个山脉数组：

    A.length >= 3
    在 0 < i < A.length - 1 条件下，存在 i 使得：
        A[0] < A[1] < ... A[i-1] < A[i]
        A[i] > A[i+1] > ... > A[A.length - 1]
 */
package _41_validMountainArray

func validMountainArray(A []int) bool {

	var firstMain int
	for i := 1; i < len(A); i++ {
		if A[i] < A[i-1] {
			firstMain = i-1
			break
		}
	}
	if firstMain == 0 {
		return false
	}
	result := true
	for i := firstMain; i < len(A)-1; i++ {
		if A[i] <= A[i+1] {
			result = false
			break
		}
	}
	return result
}

