/**
 *  @Author: JinxinZhou
 *  @Date  : 2020/10/30 10:42 上午
 */
/**
给定一个二叉树，找到最长的路径，这个路径中的每个节点具有相同值。 这条路径可以经过也可以不经过根节点。
注意：两个节点之间的路径长度由它们之间的边数表示。
*/
package _87_longestUnivaluePath

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var result int

func longestUnivaluePath(root *TreeNode) int {
	result = 0
	dfs(root)
	return result
}

func dfs(root *TreeNode) int {
	fmt.Println("ll", result)
	if root == nil {
		return 0
	}
	left := dfs(root.Left)
	right := dfs(root.Right)
	var l, r int
	if root.Left != nil && root.Val == root.Left.Val {
		l = left + 1
	}
	if root.Right != nil && root.Val == root.Right.Val {
		r = right + 1
	}
	result = max(result, l+r)
	return max(l, r)
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func TestLongestUnivaluePath(t *testing.T) {
	root := &TreeNode{Val: 1}
	l1 := &TreeNode{Val: 4}
	l2 := &TreeNode{Val: 5}
	l3 := &TreeNode{Val: 4}
	l4 := &TreeNode{Val: 4}
	l5 := &TreeNode{Val: 5}
	root.Left = l1
	root.Right = l2
	l1.Left = l3
	l1.Right = l4
	l2.Right = l5
	longestUnivaluePath(root)
}
