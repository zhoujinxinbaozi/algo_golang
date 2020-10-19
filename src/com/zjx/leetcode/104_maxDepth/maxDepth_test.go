/**
 *  @Author: JinxinZhou
 *  @Date  : 2020/10/19 12:22 下午
 */
/**
给定一个二叉树，找出其最大深度。
二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
说明: 叶子节点是指没有子节点的节点。

solution: dfs

*/
package _04_maxDepth

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var result int

func maxDepth(root *TreeNode) int {
	result = 0
	dfs(root, 0)
	return result
}

func dfs(root *TreeNode, current int) {
	if root == nil {
		return
	}

	current++

	if root.Left == nil && root.Right == nil {
		result = max(result, current)
	}

	if root.Left != nil {
		dfs(root.Left, current)
	}

	if root.Right != nil {
		dfs(root.Right, current)
	}

	current--

}

func max(i, j int) int {
	if i <= j {
		return j
	}

	return i
}
