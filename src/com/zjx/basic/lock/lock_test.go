package lock

import (
	"fmt"
	"sync"
	"testing"
)

// 互斥锁
var lock sync.Mutex
// 读写锁  可以同时读，其他的情况阻塞
// lock1.RLock()  读
// lock1.RUnlock() 读
// 写为lock
var lock1 sync.RWMutex
var count = 0
var wg sync.WaitGroup
func TestLock(t *testing.T) {

	wg.Add(2)
	go f1()
	go f1()
	wg.Wait()
	fmt.Println(count)

}

func f1(){
	defer wg.Done()
	for i:= 0; i < 500000; i ++{
		lock.Lock()
		count += 1
		lock.Unlock()
	}
}