/**
 *  @Author: JinxinZhou
 *  @Date  : 2020/10/14 9:58 上午
 */

package _00_numIslands
/**
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
岛屿总是被水包围，并且每座岛屿只能由水平方向或竖直方向上相邻的陆地连接形成。
sloution：
	使用dfs和bfs进行处理
	dfs：遍历每一个元素，使用递归的方式进行
	bfs：遍历每一个元素，需要使用队列的数据结构，将所有的可触及都加入到当中，知道队列为空即可。
 */
import (
	"fmt"
	"testing"
)

func TestNumIslands(t *testing.T) {
	arr := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	fmt.Println(numIslands(arr))
}

func numIslands(grid [][]byte) int {
	raw := len(grid)
	col := len(grid[0])
	var result int

	for i := 0; i < raw; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '0' {
				continue
			}
			//dfs(grid, i, j)
			bfs(grid, i, j)
			result++
		}
	}
	return result
}
// 深度优先 将是1 所有相邻的都设置为0   结果为有多少次深度遍历就有多少个岛屿
func dfs(grid [][]byte, x, y int) {
	raw := len(grid)
	col := len(grid[0])

	if x < 0 || x >= raw || y < 0 || y >= col || grid[x][y] == '0' {
		return
	}

	grid[x][y] = '0'

	dfs(grid, x-1, y)
	dfs(grid, x+1, y)
	dfs(grid, x, y-1)
	dfs(grid, x, y+1)

}

var queue []*Node

type Node struct {
	X int
	Y int
}

func bfs(grid [][]byte, x, y int) {

	raw := len(grid)
	col := len(grid[0])
	node := NewNode(x, y)
	queue = append(queue, node)
	for len(queue) != 0 {
		tmp := queue[0]
		x = tmp.X
		y = tmp.Y
		grid[x][y] = '0'
		queue = queue[1:]
		if judgeValid(grid, x-1, y, raw, col) {
			queue = append(queue, NewNode(x-1, y))
		}
		if judgeValid(grid, x+1, y, raw, col) {
			queue = append(queue, NewNode(x+1, y))
		}
		if judgeValid(grid, x, y-1, raw, col) {
			queue = append(queue, NewNode(x, y-1))
		}
		if judgeValid(grid, x, y+1, raw, col) {
			queue = append(queue, NewNode(x, y+1))
		}
	}
}

func NewNode(x, y int) *Node {
	return &Node{
		X: x,
		Y: y,
	}
}

func judgeValid(grid [][]byte, x, y, raw, col int) bool {
	if x >= 0 && x < raw && y >= 0 && y < col && grid[x][y] == '1' {
		return true
	}
	return false
}
