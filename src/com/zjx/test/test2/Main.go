package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

/**
 * @Auther: zhoujinxin@bytedance.com
 * @Date: 2020/5/29 11:15
 */

type Demo struct {
	ID   int
	Name string
}

func main() {
	// init
	var arrInt []*Demo
	mutex := sync.Mutex{}

	maxCnt := 16
	cnt := 0
	for i := 0; i < maxCnt; i++ {
		demo := &Demo{
			ID:   i,
			Name: strconv.Itoa(i),
		}
		arrInt = append(arrInt, demo)
	}

	ch := make(chan *Demo, 5)

	// produce
	go func() {
		for _, value := range arrInt {
			ch <- value
		}
	}()

	// consume
	for i := 0; i < 5; i++ {
		go func() {
			for {
				value, ok := <-ch
				if !ok {
					break
				}
				mutex.Lock()
				cnt += 1
				mutex.Unlock()
				change(value)
			}
		}()
	}

	for cnt != maxCnt {
		time.Sleep(time.Millisecond * 10)
	}
	close(ch)
	fmt.Println(arrInt)
}

func change(d *Demo) {
	d.Name = d.Name + "???"
	fmt.Println(d.Name)
}
