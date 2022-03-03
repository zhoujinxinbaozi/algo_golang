package atomic

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestAtomicSafe(t *testing.T) {
	var sum int64 = 0
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&sum, 1)
		}()
	}
	wg.Wait()
	fmt.Printf("sum : %v\n", sum)
}

// lock 同一时间只能处理一个事件  通过atomic实现的
func TestAtomicLock(t *testing.T) {
	ss := &s{}
	for i := 0; i < 5; i++ {
		tmp := i
		go func() {
			var isBreak bool
			for !isBreak {
				if ss.lock() {
					fmt.Println("get lock success, current handler num is ", tmp)
					time.Sleep(1 * time.Second)
					isBreak = true
					ss.unLock()
				}
			}
		}()
	}
	time.Sleep(10 * time.Second)
}

type s struct {
	Value int64
}

func (p *s) lock() bool {
	return atomic.CompareAndSwapInt64(&p.Value, 0, 1)
}

func (p *s) unLock() {
	atomic.StoreInt64(&p.Value, 0)
}
