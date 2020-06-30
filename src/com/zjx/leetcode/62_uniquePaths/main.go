package main

import "fmt"

/**
 * @Auther: zhoujinxin@bytedance.com
 * @Date: 2020/6/25 16:56
 */

/**
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。

问总共有多少条不同的路径？
*/
func main() {
	uniquePaths(7, 3)
}

func uniquePaths(m int, n int) int {
	f := make([][]int, m)
	for i := 0; i < m; i++ {
		f[i] = make([]int, n)
	}

	//for i := 0; i < m; i++ {
	//	f[i][0] = 1
	//}
	//for j := 1; j < n; j++ {
	//	f[0][j] = 1
	//}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 { // 初始化
				f[i][j] = 1
				continue
			}
			f[i][j] = f[i-1][j] + f[i][j-1] // 状态转移方程
		}
	}
	fmt.Print(f)

	return f[m-1][n-1]

}
