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

var result int
var tmp []int
var markSlice [][]int

func longestIncreasingPath(matrix [][]int) int {
	result = 0
	tmp = nil
	markSlice = nil

	if matrix == nil {
		return result
	}

	// mark is not visit
	for _, arr := range matrix {
		var tmp []int
		for i := 0; i < len(arr); i++ {
			tmp = append(tmp, 0)
		}
		markSlice = append(markSlice, tmp)
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			dfs(matrix, i, j)
		}
	}
	return result
}

func dfs(matrix [][]int, i, j int) {

	row := len(matrix)
	col := len(matrix[0])

	tmp = append(tmp, matrix[i][j])
	markSlice[i][j] = -1 // mark is visit

	var canThrough bool

	// left
	if isLegal(i-1, j, row, col) && markSlice[i-1][j] == 0 && matrix[i-1][j] > tmp[len(tmp)-1] {
		dfs(matrix, i-1, j)
		markSlice[i][j] = 0
	}

	// right
	if isLegal(i+1, j, row, col) && markSlice[i+1][j] == 0 && matrix[i+1][j] > tmp[len(tmp)-1] {
		dfs(matrix, i+1, j)
		markSlice[i][j] = 0
		canThrough = true
	}

	// up
	if isLegal(i, j-1, row, col) && markSlice[i][j-1] == 0 && matrix[i][j-1] > tmp[len(tmp)-1] {
		dfs(matrix, i, j-1)
		markSlice[i][j] = 0
		canThrough = true
	}

	// down
	if isLegal(i, j+1, row, col) && markSlice[i][j+1] == 0 && matrix[i][j+1] > tmp[len(tmp)-1] {
		dfs(matrix, i, j+1)
		markSlice[i][j] = 0
		canThrough = true
	}

	if !canThrough {
		result = max(tmp)
	}

	markSlice[i][j] = 0
	tmp = tmp[:len(tmp)-1]
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
