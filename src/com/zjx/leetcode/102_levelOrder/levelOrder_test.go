/**
 *  @Author: JinxinZhou
 *  @Date  : 2020/10/15 11:18 上午
 */

package _02_levelOrder

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	var tmp []*TreeNode
	tmp = append(tmp, root)
	result = append(result, []int{root.Val})
	length := 1
	l := 0
	for len(tmp) != 0 {
		cur := tmp[0]
		tmp = tmp[1:]
		if cur.Left != nil {
			tmp = append(tmp, cur.Left)
		}
		if cur.Right != nil {
			tmp = append(tmp, cur.Right)
		}
		if l+1 != length {
			l++
		} else {
			l = 0
			length = len(tmp)
			var c []int
			for _, v := range tmp {
				c = append(c, v.Val)
			}
			if c != nil {
				result = append(result, c)
			}
		}
	}
	return result
}

func TestLevelOrder(t *testing.T){
	t3 := TreeNode{Val: 3}
	t9 := TreeNode{Val: 9}
	t20 := TreeNode{Val: 20}
	t15 := TreeNode{Val: 15}
	t7 := TreeNode{Val: 7}
	t3.Left = &t9
	t3.Right = &t20
	t20.Left = &t15
	t20.Right = &t7

	fmt.Println(levelOrder(&t3))
}

