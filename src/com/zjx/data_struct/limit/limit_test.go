package limit

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"testing"
	"time"
)

// 场景  必须要处理的请求  不允许丢失
func TestLimitWait(t *testing.T) {
	limiter := newLimiter(100, 100)

	ctx := context.Background() // 可以通过设置ctx的超时时间等
	for i := 0; i < 10; i++ {
		tmp := i
		go func() {
			// 获取不到会等待 直到获取到足够的令牌
			err := limiter.WaitN(ctx, 100)
			if err != nil {
				fmt.Printf("current num : %v, err : %v\n", tmp, err)
				return
			}
			fmt.Printf("handler num : %v\n", tmp)
		}()
	}
	time.Sleep(10 * time.Second)
}

func newLimiter(rateLimit, cap int) *rate.Limiter {
	return rate.NewLimiter(rate.Limit(rateLimit), cap)
}

func TestLimitAllow(t *testing.T) {
	limiter := newLimiter(1, 10)
	for i := 0; i < 100; i++ {
		tmp := i
		go func() {
			time.Sleep(time.Duration(tmp) * time.Millisecond)
			allow := limiter.Allow()
			fmt.Printf("current : %v, allow : %v\n", tmp, allow)
		}()
	}
	time.Sleep(2 * time.Second)
}
