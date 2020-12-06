/**
 *  @Author: JinxinZhou
 *  @Date  : 2020/11/14 3:29 下午
 */

package max_heap

import (
	"fmt"
	"testing"
)

/**
topK 问题，使用小顶堆
*/
func TestMaxHeap(t *testing.T) {
	h := initHeap(3)
	arr := []int{3, 9, 11, 7}
	for _, v := range arr {
		h.insertData(v)
	}
	fmt.Println(h.data)
	fmt.Println(h.len)
}

type Heap struct {
	data []int
	len  int
}

func initHeap(len int) *Heap {
	return &Heap{
		data: make([]int, 0),
		len:  len,
	}
}

func (p *Heap) insertData(ele int) {
	if len(p.data) == 0 {
		p.data = append(p.data, ele)
		return
	}

	p.buildHeap(ele)

}

func (p *Heap) buildHeap(ele int) {

	if ele < p.data[0] {
		return
	}
	p.data = append(p.data, ele)
	last := len(p.data)-1

	if last & 1 == 1 {
		for last != 0 && p.data[last] < p.data[last/2] {
			p.data[last], p.data[last/2] = p.data[last/2], p.data[last]
			last = last/2
		}
	}
	if last & 1 == 1 {
		for last != 0 && p.data[last] < p.data[last/2-1] {
			p.data[last], p.data[last/2-1] = p.data[last/2-1], p.data[last]
			last = last/2-1
		}
	}
}

func TestUint8 (t *testing.T) {
	fmt.Println(float64(10)/float64(3))
}

func TestMap(t *testing.T) {
	m := make(map[string]string)
	m["1"] = "1v"
	m["2"] = "2v"
	for v := range m {
		fmt.Println(v)
	}
}