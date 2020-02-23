package main

/**
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4

solution:
	两个链表中最小的头结点放入到dummy.next中，直到一个为空，另一个全部加入即可。
 */
func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tmp := dummy
	for {
		if l1 == nil || l2 == nil {
			break
		}
		if l1.Val > l2.Val {
			tmp.Next = l2
			tmp = tmp.Next
			l2 = l2.Next
		}else{
			tmp.Next = l1
			tmp = tmp.Next
			l1 = l1.Next
		}
	}
	if l1 == nil {
		for {
			if l2 == nil {
				break
			}
			tmp.Next = l2
			tmp = tmp.Next
			l2 = l2.Next
		}
	}else {
		for {
			if l1 == nil {
				break
			}
			tmp.Next = l1
			tmp = tmp.Next
			l1 = l1.Next
		}
	}
	return dummy.Next
}
