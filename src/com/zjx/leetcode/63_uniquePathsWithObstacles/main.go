package main

/**
 * @Auther: zhoujinxin@bytedance.com
 * @Date: 2020/6/26 16:13
 */

/**
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。

现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？



网格中的障碍物和空位置分别用 1 和 0 来表示。

说明：m 和 n 的值均不超过 100。

示例 1:

输入:
[
  [0,0,0],
  [0,1,0],
  [0,0,0]
]
输出: 2
解释:
3x3 网格的正中间有一个障碍物。
从左上角到右下角一共有 2 条不同的路径：
1. 向右 -> 向右 -> 向下 -> 向下
2. 向下 -> 向下 -> 向右 -> 向右
*/
func main() {
	uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}})
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	result := make([][]int, len(obstacleGrid))

	for index, _ := range result {
		result[index] = make([]int, len(obstacleGrid[0]))
	}

	//copy(result, obstacleGrid)

	hasZA := false
	for i := 0; i < len(obstacleGrid); i++ { // 初始化列
		if obstacleGrid[i][0] == 1 {
			hasZA = true
		}
		if !hasZA {
			result[i][0] = 1
			continue
		}
		hasZA = true
		result[i][0] = 0
	}

	hasZA = false
	for j := 0; j < len(obstacleGrid[0]); j++ { // 初始化行
		if obstacleGrid[0][j] == 1 {
			hasZA = true
		}
		if !hasZA {
			result[0][j] = 1
			continue
		}
		hasZA = true
		result[0][j] = 0
	}

	for i := 1; i < len(obstacleGrid); i++ {
		for j := 1; j < len(obstacleGrid[0]); j++ {
			if obstacleGrid[i][j] == 1 {
				result[i][j] = 0
				continue
			}
			result[i][j] = result[i-1][j] + result[i][j-1]
		}
	}

	return result[len(result)-1][len(result[0])-1]

}
