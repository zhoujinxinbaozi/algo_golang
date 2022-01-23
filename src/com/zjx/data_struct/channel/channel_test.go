/**
 *  @Author: JinxinZhou
 *  @Date  : 2022/1/19 3:33 下午
 */

package channel

import (
	"fmt"
	"testing"
)

/*
关闭未初始化的channle(nil)会panic
重复关闭同一channel会panic
向以关闭channel发送消息会panic
从已关闭channel读取数据，不会panic，若存在数据，则可以读出未被读取的消息，若已被读出，则获取的数据为零值，可以通过ok-idiom的方式，判断channel是否关闭
channel的关闭操作，会产生广播消息，所有向channel读取消息的goroutine都会接受到消息
*/

// TestChannelClose
// 在channel关闭时，所有向channe读取消息的goroutinue都会接收到消息  值为对应类型的0值
func TestChannelClose(t *testing.T) {
	done := make(chan int)
	close(done)
	select {
	case <-done:
		fmt.Println("receive close singal")
	default:
		fmt.Println("normal quite")
	}
}

func TestChannel(t *testing.T) {
	// 初始化channel
	ch := make(chan int, 3)
	// 发送消息
	ch <- 1
	ch <- 2
	close(ch)

	// 只会打印1和2 ch关闭后 range只能读取到之前写入的数据
	for c := range ch {
		fmt.Println(c)
	}

	// 会打印 1 2 和 无限个0
	for {
		select {
		case i := <-ch:
			fmt.Println("i : ", i)
		}
	}

	// 只会打印1和2
	quite := false
	for {
		if quite {
			break
		}
		select {
		case i, ok := <-ch:
			if !ok {
				quite = true
				break // 跳出select 后面的语句不在执行
			}
			fmt.Println("i : ", i)
		}
	}
}
