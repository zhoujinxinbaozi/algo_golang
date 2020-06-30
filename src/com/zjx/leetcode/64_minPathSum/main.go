package main

/**
 * @Auther: zhoujinxin@bytedance.com
 * @Date: 2020/6/26 16:46
 */

func main() {
	minPathSum([][]int{{1, 2, 5}, {3, 2, 1}})
}

func minPathSum(grid [][]int) int {
	result := make([][]int, len(grid))

	for index, _ := range result {
		result[index] = make([]int, len(grid[0]))
	}
	result[0][0] = grid[0][0]
	for i := 1; i < len(grid); i++ {
		result[i][0] = result[i-1][0] + grid[i][0]
	}

	for j := 1; j < len(grid[0]); j++ {
		result[0][j] = result[0][j-1] + grid[0][j]
	}

	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			min := GetMin(result[i][j-1], result[i-1][j])
			result[i][j] = min + grid[i][j]
		}
	}
	//fmt.Println(result)
	return result[len(result)-1][len(result[0])-1]
}
func GetMin(i, j int) int {
	if i < j {
		return i
	}
	return j
}
