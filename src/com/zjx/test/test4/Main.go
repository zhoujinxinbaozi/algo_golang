package main

import (
	"fmt"
	"sync"
	"time"
)

/**
 * @Auther: zhoujinxin@bytedance.com
 * @Date: 2020/6/1 14:44
 */

func main() {
	//wgImpl()
	ChanImpl()

}

func wgImpl() {
	cnt := 0
	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}
	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			mutex.Lock()
			cnt++
			mutex.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("cnt : ", cnt)
}

func ChanImpl() {
	cnt := 0
	var mutex sync.Mutex
	ch := make(chan int, 5)
	go func() {
		for i := 0; i < 1000; i++ {
			ch <- i
		}
	}()

	for i := 0; i < 5; i++ {
		go func() {
			for {
				value, ok := <-ch
				if !ok {
					return
				}
				mutex.Lock()
				cnt++
				mutex.Unlock()
				fmt.Println(value)
			}
		}()
	}

	if cnt != 1000 {
		time.Sleep(time.Millisecond * 100)
	}
	close(ch)
}
