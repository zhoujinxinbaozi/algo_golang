/**
 *  @Author: JinxinZhou
 *  @Date  : 2021/5/31 3:39 下午
 */

package interview

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// hasCycle 判断链表是否有环
// 慢指针走一步  快指针走一步
func hasCycle(head *ListNode) bool {
	if head != nil {
		slow, fast := head, head
		for fast != nil && fast.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
			if slow == fast {
				return true
			}
		}
	}
	return false
}

func TestHasCycle(t *testing.T)  {
	head := &ListNode{Val: 0}
	h1 := &ListNode{Val: 0}
	h2 := &ListNode{Val: 0}
	h3 := &ListNode{Val: 0}
	h4 := &ListNode{Val: 0}
	head.Next = h1
	h1.Next = h2
	h2.Next = h3
	h3.Next = h4
	h4.Next = h2
	fmt.Println(hasCycle(head))
}