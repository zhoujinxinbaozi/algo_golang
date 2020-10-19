/**
 *  @Author: JinxinZhou
 *  @Date  : 2020/10/19 2:56 下午
 */
/**
测试不通过
 */
package _29_longestIncreasingPath

import (
	"fmt"
	"testing"
)

func TestLongestIncreasingPath(t *testing.T) {
	nums := [][]int{
		{7, 8, 9},
		{9, 7, 6},
		{7, 2, 3},
	}
	fmt.Println(longestIncreasingPath(nums))
}

type Node struct {
	I int
	J int
}

var result int
var tmp []int
var markNodes []*Node

func longestIncreasingPath(matrix [][]int) int {
	result = 0
	tmp = nil

	if matrix == nil {
		return result
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			markNodes = nil
			dfs(matrix, i, j)
		}
	}
	return result
}

func dfs(matrix [][]int, i, j int) {

	row := len(matrix)
	col := len(matrix[0])

	if !isLegal(i, j, row, col) {
		return
	}

	tmp = append(tmp, matrix[i][j])

	var canThrough bool

	if !isThrough(i, j) {
		markNodes = append(markNodes, &Node{
			I: i,
			J: j,
		})
	}

	if isLegal(i-1, j, row, col) && !isThrough(i-1, j) && matrix[i-1][j] > tmp[len(tmp)-1] {
		canThrough = true
		markNodes = append(markNodes, &Node{
			I: i - 1,
			J: j,
		})
		dfs(matrix, i-1, j)
	}

	if isLegal(i+1, j, row, col) && !isThrough(i+1, j) && matrix[i+1][j] > tmp[len(tmp)-1] {
		canThrough = true
		markNodes = append(markNodes, &Node{
			I: i + 1,
			J: j,
		})
		dfs(matrix, i+1, j)
	}

	if isLegal(i, j-1, row, col) && !isThrough(i, j-1) && matrix[i][j-1] > tmp[len(tmp)-1] {
		canThrough = true
		markNodes = append(markNodes, &Node{
			I: i,
			J: j - 1,
		})
		dfs(matrix, i, j-1)
	}

	if isLegal(i, j+1, row, col) && !isThrough(i, j+1) && matrix[i][j+1] > tmp[len(tmp)-1] {
		canThrough = true
		markNodes = append(markNodes, &Node{
			I: i,
			J: j + 1,
		})
		dfs(matrix, i, j+1)
	}

	if !canThrough {
		fmt.Println(tmp)
		result = max(tmp)
	}

	tmp = tmp[:len(tmp)-1]
}

func isThrough(i, j int) bool {
	for _, node := range markNodes {
		if node.I == i && node.J == j {
			return true
		}
	}
	return false
}

func max(tmp []int) int {
	if len(tmp) > result {
		return len(tmp)
	}
	return result
}

func isLegal(i, j, row, col int) bool {
	if i >= 0 && i < row && j >= 0 && j < col {
		return true
	}
	return false
}
