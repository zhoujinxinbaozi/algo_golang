/**
 *  @Author: JinxinZhou
 *  @Date  : 2022/1/19 1:24 下午
 */

package priority_queue

import (
	"fmt"
	"testing"
)

// 轮询事件
type PollEvent struct {
	UserList int
	Priority uint8
}

type PollQueue struct {
	ChanList []chan *PollEvent
}

// TestQueue
func TestQueue(t *testing.T) {
	pollQueue := &PollQueue{
		ChanList: make([]chan *PollEvent, 3),
	}
	pollQueue.ChanList[0] = make(chan *PollEvent, 50)
	pollQueue.ChanList[1] = make(chan *PollEvent, 50)
	pollQueue.ChanList[2] = make(chan *PollEvent, 50)
	go pollQueue.produce()
	pollQueue.consume()
}

func (p *PollQueue) consume() {
	for {
		get := false
		for _, queue := range p.ChanList {
			if get {
				break
			}
			select {
			case s := <-queue:
				fmt.Printf("user : %v, pri : %v\n", s.UserList, s.Priority)
				get = true
			default:
			}
		}
	}
}

// index channel的数组下标
// start 起始数字
// end   结束数字
// priority 优先级
func (p *PollQueue) Add(index, start, end int, priority uint8) {
	for i := start; i < end; i++ {
		p.ChanList[index] <- &PollEvent{i, priority}
	}
}

func (p *PollQueue) produce() {
	go p.Add(0, 0, 100, 3)
	go p.Add(1, 100, 200, 2)
	go p.Add(2, 200, 300, 1)
}
