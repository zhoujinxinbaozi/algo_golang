/**
 *  @Author: JinxinZhou
 *  @Date  : 2021/9/6 10:46 上午
 */

package project

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// 模拟超时返回的代码
func TestTimeout(t *testing.T) {
	ch := make(chan string)
	quit := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- strconv.Itoa(i)
			//time.Sleep(time.Second)
		}
		close(ch)
	}()

	go func() {
		timeout := time.After(time.Second * 2) // 注意，定时器必须放在select外面
		for {
			select {
			case <-timeout:
				fmt.Println("2 second timout,then send one message to channel quit ,and exit for loop and goroutine")
				quit <- true
				return
			case str, isClosed := <-ch:
				if !isClosed {
					quit <- true
					return
				}
				fmt.Println("string: ", str)
			}
		}
	}()
	<-quit
	close(quit)
	fmt.Println("endof process")
}
