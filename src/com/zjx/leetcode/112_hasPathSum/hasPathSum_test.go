/**
 *  @Author: JinxinZhou
 *  @Date  : 2020/10/14 2:37 下午
 */
/**
给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。
说明: 叶子节点是指没有子节点的节点。
solution:
dfs
 */
package _12_hasPathSum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var result bool
func hasPathSum(root *TreeNode, sum int) bool {
	result = false
	if root != nil {
		dfs(root, 0, sum)
	}
	return result
}

func dfs(root *TreeNode, target, sum int) {
	if target+root.Val == sum && root.Left == nil && root.Right == nil {
		result = true
		return
	}

	if root.Left != nil {
		dfs(root.Left, target+root.Val, sum)
	}

	if root.Right != nil {
		dfs(root.Right, target+root.Val, sum)
	}
}