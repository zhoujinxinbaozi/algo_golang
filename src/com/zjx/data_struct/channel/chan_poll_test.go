/**
 *  @Author: JinxinZhou
 *  @Date  : 2022/1/26 1:54 下午
 */

package channel

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
	"testing"
	"time"
)

var commercialPushthreadPool *ants.Pool

func TestPoll(t *testing.T) {
	commercialPushthreadPool, _ = ants.NewPool(100)
	wg := sync.WaitGroup{}
	length := 10000
	wg.Add(length)
	now := time.Now()
	for i:= 0; i < length; i++ {
		tmp := i
		commercialPushthreadPool.Submit(func() {
			defer wg.Done()
			time.Sleep(time.Duration(tmp) * time.Microsecond)
		})
	}
	wg.Wait()
	fmt.Println("success, cost : ", time.Since(now))
}
