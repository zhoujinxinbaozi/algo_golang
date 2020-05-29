package main

import (
	"fmt"
	"time"
)

/**
 * @Auther: zhoujinxin@bytedance.com
 * @Date: 2020/5/29 11:15
 */

func main() {
	// init
	var arrInt []int

	maxCnt := 16
	cnt := 0
	for i := 0; i < maxCnt; i++ {
		arrInt = append(arrInt, i)
	}

	ch := make(chan int, 5)

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
				cnt += 1
				fmt.Println(value)
			}
		}()
	}

	for cnt != maxCnt {
		time.Sleep(time.Millisecond * 10)
	}
	close(ch)
}
