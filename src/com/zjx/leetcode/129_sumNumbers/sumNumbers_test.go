/**
 *  @Author: JinxinZhou
 *  @Date  : 2020/10/29 7:35 下午
 */

package _29_sumNumbers

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var tmp []int
var result int

func TestSumNumbers(t *testing.T) {
	root := &TreeNode{Val: 1}
	l := &TreeNode{Val: 2}
	r := &TreeNode{Val: 3}
	root.Left = l
	root.Right = r
	fmt.Print(sumNumbers(root))
}

func sumNumbers(root *TreeNode) int {
	result = 0
	tmp = nil
	dfs(root)
	return result
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	tmp = append(tmp, root.Val)
	if root.Left == nil && root.Right == nil {
		cal()
		fmt.Println(tmp)
	}

	if root.Left != nil {
		dfs(root.Left)
	}
	if root.Right != nil {
		dfs(root.Right)
	}
	tmp = tmp[:len(tmp)-1]
}

func cal() {
	var r int
	for i, v := range tmp {
		if i == 0 {
			r = v
		} else {
			r = r*10 + v
		}
	}
	result += r
}
