package observe

import (
	"fmt"
	"sync"
	"testing"
)

// Subject
type Subject struct {
	ObserverList []Observer
	Lock         sync.Mutex
}

func (s *Subject) doSomething() {
	for _, v := range s.ObserverList {
		v.update()
	}
}

// AddObserver 添加观察者
func (s *Subject) AddObserver(o Observer) {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	s.ObserverList = append(s.ObserverList, o)
}

// Observer 观察者接口
type Observer interface {
	update()
}

// ObserverImp 观察者实现者
type ObserverImp struct {
	id int
}

// update 观察者在接到通知的操作
func (o *ObserverImp) update() {
	fmt.Printf("id : %+v, update\n", o.id)
}

type ObserverList []Observer

// TestObserve 观察者测试
func TestObserve(t *testing.T) {
	var s *Subject

	var observerList ObserverList
	observerList = make([]Observer, 0)
	var lock sync.Mutex

	// 创建主题
	s = &Subject{
		ObserverList: observerList,
		Lock:         lock,
	}

	var wg sync.WaitGroup
	// 多协成 添加观察者
	for i := 0; i < 1000; i++ {
		o := &ObserverImp{
			id: i,
		}
		go func() {
			wg.Add(1)
			defer wg.Done()
			go s.AddObserver(o)
		}()
	}

	// make sure observer add to list
	wg.Wait()

	// simulation change
	s.doSomething()
	fmt.Printf("len : %+v", len(s.ObserverList))
}
