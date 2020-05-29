package main

import (
	"fmt"
	"sync"
	"time"
)

/**
 * @Auther: zhoujinxin@bytedance.com
 * @Date: 2020/5/29 10:48
 */

func main() {
	// init
	var arrInt []int
	for i := 0; i < 16; i++ {
		arrInt = append(arrInt, i)
	}

	// separate
	var result [][]int
	begin := 0
	for {
		if begin >= len(arrInt) {
			break
		}
		tmp := arrInt[begin:findMin(len(arrInt), begin+5)]
		begin += 5
		result = append(result, tmp)
	}

	// handle
	var wg sync.WaitGroup
	for _, arr := range result {
		wg.Add(len(arr))
		for _, value := range arr {
			value1 := value
			go func() {
				defer wg.Done()
				fmt.Println(value1)
				time.Sleep(time.Second)
			}()
		}
		wg.Wait()
		fmt.Println("===============")
	}

}

func findMin(i, j int) int {
	if i < j {
		return i
	}
	return j
}
