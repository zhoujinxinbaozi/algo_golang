package interview

import (
	"fmt"
	"testing"
)

func TestRevert(t *testing.T) {
	head := init1()
	printf(*head,"before revert", "before revert")
	head = revert(head)
	printf(*head,"after revert", "after revert")
}

func revert(head *Node) *Node {
	dunmmy := Node{
		value: -1,
	}
	dunmmy.next = head
	prev := dunmmy.next
	cur := prev.next
	for {
		if cur == nil {
			break
		}
		prev.next = cur.next
		cur.next = dunmmy.next
		dunmmy.next = cur
		cur = prev.next
	}
	return dunmmy.next
}

func printf(head Node, begin, end string) {
	fmt.Println("=====", begin, "=====")
	h := &head
	for {
		if h == nil {
			return
		}
		fmt.Println(h.value)
		h = h.next
	}
}

func init1() *Node {
	node0 := Node{
		value: 0,
	}
	node1 := Node{
		value: 1,
	}
	node2 := Node{
		value: 2,
	}
	node0.next = &node1
	node1.next = &node2
	return &node0
}

type Node struct {
	value int
	next  *Node
}
