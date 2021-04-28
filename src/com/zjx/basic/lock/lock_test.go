package lock

import (
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/robfig/cron"
)

// 互斥锁
var lock sync.Mutex

// 读写锁  可以同时读，其他的情况阻塞
// lock1.RLock()  读
// lock1.RUnlock() 读
// 写为lock
var lock1 sync.RWMutex
var count = 0
var count1 = 0
var wg sync.WaitGroup

func TestLock(t *testing.T) {
	cur := time.Now()
	wg.Add(2)
	go f1()
	go f1()
	wg.Wait()
	fmt.Println(count)
	fmt.Println("lock cost : ", time.Now().Sub(cur).Milliseconds())
}

type method func()

func TestUnLock(t *testing.T) {
	cur := time.Now()
	ch := make(chan method, 2)
	ch <- f2
	ch <- f2
	close(ch)

	for f := range ch {
		f()
	}
	fmt.Println("count : ", count)
	fmt.Println("cost :", time.Now().Sub(cur).Milliseconds())
}

func f2() {

	for i := 0; i < 500000000; i++ {
		count1 += 1
	}
}

func f1() {
	defer wg.Done()

	lock.Lock()
	for i := 0; i < 500000000; i++ {
		count += 1
	}
	lock.Unlock()
}

// TestTimeLocker
func TestTimeLocker(t *testing.T) {
	timeLocker := NewTimeLocker()
	fmt.Printf("begin : %v, end : %v\n", timeLocker.begin, timeLocker.end)
	timeLocker.LockAfter(time.Now().Unix(), 30)
	fmt.Printf("begin : %v, end : %v\n", timeLocker.begin, timeLocker.end)

	time.Sleep(time.Duration(1) * time.Second)

	timeLocker.LockAfter(time.Now().Unix(), 30)
	fmt.Printf("begin : %v, end : %v\n", timeLocker.begin, timeLocker.end)
}

// TestSplit
func TestSplit(t *testing.T) {
	str := "1	2			3"
	arr := strings.Split(str, "\t")
	fmt.Println("len = ", len(arr), " arr = ", arr)
}

// TestTime
func TestTime(t *testing.T) {
	now := time.Now().Format("20060102")
	nowLocal := time.Now().Local().Format("20060102")
	fmt.Println("now : ", now)
	fmt.Println("nowLocal : ", nowLocal)
	fmt.Println(time.Now().Unix())
}

// ScheduledDelete
func TestScheduledDelete(t *testing.T) {
	c := cron.New()
	spec := "00 29 16 * * ?"
	c.AddFunc(spec, func() {
		fmt.Println("123")
	})
	c.Start()
	select{}
}