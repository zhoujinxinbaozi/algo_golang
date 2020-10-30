/**
 *  @Author: JinxinZhou
 *  @Date  : 2020/10/28 2:21 下午
 */

package _202_kthToLast

import "testing"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
倒数第k个节点
 */

func kthToLast(head *ListNode, k int) int {

	tmp := head
	var l int

	for tmp != nil {
		l++
		tmp = tmp.Next
	}

	for l-k != 0 {
		head = head.Next
		k++
	}

	return head.Val
}

func TestKthToLast(t *testing.T) {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	kthToLast(n1, 2)
}
