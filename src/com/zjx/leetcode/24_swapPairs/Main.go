package main

import "fmt"
/**
	problem:
	给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
	solution:

 */
func main() {
	l1 := &ListNode{Val:1,}
	l2 := &ListNode{Val:2,}
	l3 := &ListNode{Val:3,}
	l4 := &ListNode{Val:4,}
	l2.Next = l1
	l1.Next = l3
	l3.Next = l4
	swapPairs(l2)

}

func print (result *ListNode){
	for {
		if result == nil {
			break
		}
		fmt.Print(result.Val)
		result = result.Next
	}
	fmt.Println()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	index := 1
	dummy1 := &ListNode{Val:-1}
	dummy2 := &ListNode{Val:-1}
	t1 := dummy1
	t2 := dummy2
	for {
		if head == nil {
			break
		}
		if index %2 == 0 {
			t2.Next = head
			t2 = t2.Next
			head = head.Next
			index ++
			print(dummy2.Next)
		}else {
			t1.Next = head
			t1 = t1.Next
			head = head.Next
			index ++
			print(dummy1.Next)
		}
	}



	result := &ListNode{}
	tmp := result
	dummy2 = dummy2.Next
	dummy1 = dummy1.Next
	for {
		if dummy2 == nil && dummy2 == nil {
			break
		}
		if dummy2 != nil {
			tmp.Next = dummy2
			dummy2 = dummy2.Next
		}else {
			tmp.Next = dummy1.Next
			dummy1 = dummy1.Next
		}
	}
	return result.Next
}
