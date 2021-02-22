/**
 *  @Author: JinxinZhou
 *  @Date  : 2021/2/23 11:41 上午
 */

package panic

import (
	"fmt"
	"testing"
	"time"
)

// TestPanic panic 只会触发当前 Goroutine 的 defer
func TestPanic(t *testing.T) {
	defer func() {
		fmt.Println("TestPanic Defer")
	}()

	go func() {
		defer func() {
			fmt.Println("Goroutine wide Defer")
			defer func() {
				fmt.Println("Goroutine inner Defer")
			}()
		}()
		panic("error")
	}()

	time.Sleep(time.Second * 1)
}
