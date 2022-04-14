package task_pool

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"
)

// -cpu指定 go test -bench=BenchmarkDemo com/zjx/algo_golang/src/com/zjx/data_struct/task_pool/ -v -cpu=1

// go test -bench=BenchmarkDemo com/zjx/algo_golang/src/com/zjx/data_struct/task_pool/ -v
func BenchmarkDemo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Sleep(time.Duration(i) * time.Millisecond)
	}
}

type Mark struct {
	Name string
	Pass bool
}

// BenchmarkTaskPool
func BenchmarkTaskPool(b *testing.B) {
	ch := make(chan *Mark)
	// 该协程 异步调用多个RPC 通过ch进行通信
	go func() {
		wg := sync.WaitGroup{}
		// 模拟taskpool
		for i := 0; i <= b.N; i++ {
			data := i
			wg.Add(1)
			go func() {
				defer wg.Done()
				ch <- Rpc(data)
			}()
		}
		wg.Wait()
		close(ch)
	}()

	// 通过channel close 来处理收到的数据
	for m := range ch {
		fmt.Printf("handle data : %v\n", m)
	}
}

// Rpc 模拟rpc调用
// @param data
// @return *Mark
func Rpc(data int) *Mark {
	result := &Mark{}

	time.Sleep(time.Duration(10) * time.Millisecond)
	if data&1 == 1 {
		result.Pass = true
		result.Name = strconv.Itoa(data)
	} else {
		result.Pass = false
		result.Name = strconv.Itoa(data)
	}
	return result
}

// go test -bench='BenchmarkGen*' com/zjx/algo_golang/src/com/zjx/data_struct/task_pool/ -v -benchmem
func BenchmarkGenerateWithCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateWithCap(10000)
	}
}

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generate(10000)
	}
}
func generateWithCap(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func generate(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}
