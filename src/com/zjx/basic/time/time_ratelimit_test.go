/**
 *  @Author: JinxinZhou
 *  @Date  : 2020/9/27 4:03 下午
 */

package time

import (
	"fmt"
	"go.uber.org/ratelimit"
	"testing"
)

// TestRateLimit
func TestRateLimit(t *testing.T) {
	// every minute get request
	limit := ratelimit.New(1)
	for {
		// block take
		limit.Take()
		fmt.Println(1)
	}
}
