/**
 *  @Author: JinxinZhou
 *  @Date  : 2021/2/18 3:49 下午
 */

package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// TestContextWithCancel
func TestContextWithCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go Watch(ctx, i)
	}
	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

// Watch
func Watch(ctx context.Context, str int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("exit watch name : ", str)
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("monitor name : ", str)
		}
	}
}

// TestContextWithValue
func TestContextWithValue(t *testing.T) {
	name := "name"
	val := "value"
	ctx := context.WithValue(context.Background(), name, val)
	GetValue(ctx, name)
	time.Sleep(1 * time.Second)
}

// GetValue
func GetValue(ctx context.Context, key string) {
	fmt.Println(ctx.Deadline())
	fmt.Println(ctx.Value(key))
}
