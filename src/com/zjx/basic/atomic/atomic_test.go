package atomic

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

/**
 * @Auther: zhoujinxin@bytedance.com
 * @Date: 2020/8/18 21:15
 */

// 多线程计数，会出现竞争
func TestCount(t *testing.T) {
	var count int32
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			count++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("count = ", count)
}

// 多线程计数，使用原子操作
func TestAtomicCount(t *testing.T) {
	var count int32
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&count, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("count = ", atomic.LoadInt32(&count))

	var n int64
	n = 4

	fmt.Println(n)
	fmt.Println(^(n - 1))
	fmt.Println(^n)
}
