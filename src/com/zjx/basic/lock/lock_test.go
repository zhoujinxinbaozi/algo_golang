package lock

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 互斥锁
var lock sync.Mutex

// 读写锁  可以同时读，其他的情况阻塞
// lock1.RLock()  读
// lock1.RUnlock() 读
// 写为lock
var lock1 sync.RWMutex
var count = 0
var count1 = 0
var wg sync.WaitGroup

func TestLock(t *testing.T) {
	cur := time.Now()
	wg.Add(2)
	go f1()
	go f1()
	wg.Wait()
	fmt.Println(count)
	fmt.Println("lock cost : ", time.Now().Sub(cur).Milliseconds())
}

type method func()

func TestUnLock(t *testing.T) {
	cur := time.Now()
	ch := make(chan method, 2)
	ch <- f2
	ch <- f2
	close(ch)

	for f := range ch {
		f()
	}
	fmt.Println("count : ", count)
	fmt.Println("cost :", time.Now().Sub(cur).Milliseconds())
}

func f2() {

	for i := 0; i < 500000000; i++ {
		count1 += 1
	}
}

func f1() {
	defer wg.Done()

	lock.Lock()
	for i := 0; i < 500000000; i++ {
		count += 1
	}
	lock.Unlock()
}
