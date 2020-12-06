/**
 *  @Author: JinxinZhou
 *  @Date  : 2020/12/6 12:55 下午
 */

package main

import (
	"fmt"
	"testing"
	"time"
)

func TestChan(t *testing.T) {
	var chanList []chan interface{}
	chanList = make([]chan interface{}, 0, 5)
	go func() {
		var c chan interface{}
		c = make(chan interface{}, 1)
		for i := 0; i < 10; i++ {
			var tmp = i
			go func() {
				c <- tmp
			}()
		}
		chanList = append(chanList, c)
	}()

	time.Sleep(time.Duration(100) * time.Millisecond)
	fmt.Println("==========", len(chanList), "==========")
	go func() {
		for _, v := range chanList {
			go func() {
				for {
					select {
					case vv := <-v:
						fmt.Println(vv)
					}
				}
			}()

		}
	}()

	time.Sleep(time.Duration(1) * time.Second)
}
