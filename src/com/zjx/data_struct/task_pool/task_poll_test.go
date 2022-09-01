package task_pool

import (
	"bufio"
	"context"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"testing"
	"time"
	"unsafe"
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

func TestSize(t *testing.T) {
	i := int64(2)
	fmt.Println(unsafe.Sizeof(i))
}

func TestContext(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "1", 123)
	ctx = context.WithValue(ctx, "1", 345)
	result, ok := ctx.Value("1").(int64)
	if !ok {
		fmt.Printf("failed, key : %p", ctx.Value("1"))
		return
	}
	fmt.Println(result)
}

func TestMD5(t *testing.T) {

	fi, err := os.Open("/Users/bytedance/Desktop/wikcnRf4qC6ji1fU2emfrsqtHCg")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	f, err := os.Create("/Users/bytedance/Desktop/aaa")
	if err != nil {
		fmt.Printf("create map file error: %v\n", err)
		panic("123")
	}
	defer f.Close()

	w := bufio.NewWriter(f)

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		line := string(a)
		ss := MD5(fmt.Sprintf("%v%v", line, "42b91e"))
		ss = SHA1(fmt.Sprintf("%v%v", "08a441", ss))
		fmt.Println(line, "\t", ss)
		fmt.Fprintln(w, line, "\t", ss)
	}
	w.Flush()
}

func SHA1(s string) string {

	o := sha1.New()

	o.Write([]byte(s))

	return hex.EncodeToString(o.Sum(nil))

}

func MD5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}
