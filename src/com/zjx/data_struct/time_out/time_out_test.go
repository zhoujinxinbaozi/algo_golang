/**
 *  @Author: JinxinZhou
 *  @Date  : 2022/1/19 1:27 下午
 */

package time_out

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimeOut(t *testing.T){
	wg := sync.WaitGroup{}
	start := 0
	end := 10
	wg.Add(10)
	for i := start; i < end; i++ {
		tmp := i
		go func() {
			getData(tmp)
			wg.Done()
		}()
	}
	wg.Wait()
}

// getData 模拟请求调用接口获取数据
func getData(param int) {
	chanFlag := make(chan struct{})
	timeTicker := time.NewTicker(time.Duration(1) * time.Second)
	result := 0
	go handleData(chanFlag, param, &result)
	select {
	case <- chanFlag:
		fmt.Printf("param : %v, result : %v\n", param, result)
	case <-timeTicker.C:
		fmt.Printf("param : %v is time out, result is default : %v\n", param, result)
	}
}

// handleData 处理数据
func handleData(chanFlag chan struct{}, param int, result *int) {
	timeOut := param % 2
	time.Sleep(time.Second * time.Duration(timeOut))
	chanFlag <- struct{}{}
	*result = param
}