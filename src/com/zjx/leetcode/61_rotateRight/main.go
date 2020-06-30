package main

import "fmt"

/**
 * @Auther: zhoujinxin@bytedance.com
 * @Date: 2020/6/25 16:06
 */
/**
给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。

示例 1:

输入: 1->2->3->4->5->NULL, k = 2
输出: 4->5->1->2->3->NULL
解释:
向右旋转 1 步: 5->1->2->3->4->NULL
向右旋转 2 步: 4->5->1->2->3->NULL
示例 2:

输入: 0->1->2->NULL, k = 4
输出: 2->0->1->NULL
解释:
向右旋转 1 步: 2->0->1->NULL
向右旋转 2 步: 1->2->0->NULL
向右旋转 3 步: 0->1->2->NULL
向右旋转 4 步: 2->0->1->NULL
*/
func main() {
	h3 := ListNode{
		Val:  4,
		Next: nil,
	}

	h2 := ListNode{
		Val:  3,
		Next: &h3,
	}

	h1 := ListNode{
		Val:  2,
		Next: &h2,
	}

	head := ListNode{
		Val:  1,
		Next: &h1,
	}

	r := rotateRight(&head, 2)
	fmt.Println(r)

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	if k == 0 {
		return head
	}

	headNode := head
	count := 1
	for head.Next != nil { // 收尾相连成环
		head = head.Next
		count++
	}
	head.Next = headNode

	k = k % count
	for count-k != 1 { // 走k步，为count-k 为头结点  -> 少走一步，去next为head，并断开环
		headNode = headNode.Next
		count--
	}
	result := headNode.Next
	headNode.Next = nil
	return result
}
