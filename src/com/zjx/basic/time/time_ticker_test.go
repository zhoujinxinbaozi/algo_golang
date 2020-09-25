/**
 *  @Author: JinxinZhou
 *  @Date  : 2020/9/25 2:07 下午
 */

package time

import (
	"fmt"
	"testing"
	"time"
)

// TestTimeTicker
func TestTimeTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	for {
		<-ticker.C
		fmt.Println(1111)
	}
}
