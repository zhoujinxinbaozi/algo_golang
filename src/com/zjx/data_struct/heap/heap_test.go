/**
 *  @Author: JinxinZhou
 *  @Date  : 2022/1/25 11:28 上午
 */

package heap

import (
	"container/heap"
	"fmt"
	"testing"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j] // 小于是小顶堆   大于是大顶堆
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	fmt.Printf("    %v\n", old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func TestHeap(t *testing.T) {
	h := &IntHeap{2, 1, 5, 100, 3, 6, 4, 5}
	heap.Init(h)
	heap.Push(h, 3)
	heap.Push(h, 30)
	heap.Push(h, 9)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}
