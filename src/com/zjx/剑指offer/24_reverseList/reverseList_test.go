/**
 *  @Author: JinxinZhou
 *  @Date  : 2020/11/6 11:40 上午
 */
/**
定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。
 */
package _4_reverseList

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{}
	dummy.Next = head
	prev := head
	pcur := head.Next

	for pcur != nil {
		prev.Next = pcur.Next
		pcur.Next = dummy.Next
		dummy.Next = pcur
		pcur = prev.Next
	}
	return dummy.Next
}
