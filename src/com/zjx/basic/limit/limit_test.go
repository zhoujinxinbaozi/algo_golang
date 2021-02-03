/**
 *  @Author: JinxinZhou
 *  @Date  : 2021/1/28 3:57 下午
 */

package limit

import (
	"fmt"
	"go.uber.org/ratelimit"
	"golang.org/x/time/rate"
	"sync"
	"testing"
	"time"
)

var limit ratelimit.Limiter
var lim *rate.Limiter

func Init(speed int) {
	limit = ratelimit.New(speed)
	lim = rate.NewLimiter(rate.Limit(speed), speed)
}

func TestLimit(t *testing.T) {
	Init(1)
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		temp := i
		go func() {
			before := time.Now()
			limit.Take()
			f()
			fmt.Println("i : ", temp, " cost : ", time.Now().Sub(before).Milliseconds())
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestLimit1(t *testing.T) {
	Init(1)
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		temp := i
		go func() {
			before := time.Now()
			if lim.Allow() {
				f()
				fmt.Println("i : ", temp, " cost : ", time.Now().Sub(before).Milliseconds())
			} else {
				fmt.Println("i : ", temp, "can't get token")
			}
			wg.Done()
		}()
	}
	wg.Wait()

}

func f() {

}