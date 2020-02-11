package main

/**
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？

solution:
	双指针法
*/
func main() {
	l1 := ListNode{
		Val:  0,
		Next: nil,
	}
	l2 := ListNode{
		Val:  0,
		Next: nil,
	}
	l1.Next = &l2
	removeNthFromEnd(&l1,1)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	sentinel := &ListNode{Next: head}
	fast, slow, index := sentinel, sentinel, 0
	for index < n+1 {
		fast = fast.Next
		index++
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return sentinel.Next
}

