/**
 *  @Author: JinxinZhou
 *  @Date  : 2020/12/10 8:18 下午
 */

/*
判断一颗树是否为平衡二叉树（根节点左右节点不超过一）
*/
package my

import (
	"fmt"
	"testing"
)

type Node struct {
	V     int
	left  *Node
	right *Node
}

func NewNode(v int) *Node {
	return &Node{
		V: v,
	}
}

// AddNode direction  1  左子树   2  右子树
func AddNode(root *Node, direction int, v int) *Node {
	node := NewNode(v)
	if direction == 1 {
		root.left = node
	} else {
		root.right = node
	}
	return node
}

var result bool

func TestBinaryTree(t *testing.T) {
	result = true
	root := NewNode(6)
	AddNode(root, 2, 9)
	n1 := AddNode(root, 1, 6)
	n2 := AddNode(n1, 2, 7)
	AddNode(n2, 1, 8)
	judgeIsBinaryBalanceTree(root)
	fmt.Println(result)
}

func judgeIsBinaryBalanceTree(root *Node) int {
	if root == nil {
		return 0
	}

	l := judgeIsBinaryBalanceTree(root.left)  // 左子树最大高度
	r := judgeIsBinaryBalanceTree(root.right) // 右子树最大高度

	if l-r > 1 || l-r < -1 {
		result = false
	}

	if l < r {
		return r + 1
	} else {
		return l + 1
	}
}
