/**
 *  @Author: JinxinZhou
 *  @Date  : 2021/9/6 3:14 下午
 */

package interview

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

// isSearchBinaryTree
func isSearchBinaryTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.left != nil {
		if root.left.value > root.value {
			return false
		}
	}
	if root.right != nil {
		if root.right.value < root.value {
			return false
		}
	}
	return isSearchBinaryTree(root.left) && isSearchBinaryTree(root.right)
}

func TestIsSearchBinaryTree(t *testing.T) {
	root := &TreeNode{value: 8}

	r := &TreeNode{value: 15}
	rr := &TreeNode{value: 19}
	//rl := &Node{value: 12}

	l := &TreeNode{value: 5}
	ll := &TreeNode{value: 3}
	lr := &TreeNode{value: 7}

	root.left = l
	root.right = r

	l.left = ll
	l.right = lr

	//r.left = rl
	r.right = rr

	fmt.Println(isSearchBinaryTree(root))
}
