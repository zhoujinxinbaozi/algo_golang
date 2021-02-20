/**
 *  @Author: JinxinZhou
 *  @Date  : 2021/2/18 5:57 下午
 */

package channel

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	// ...
	const MaxRandomNumber = 100000
	const NumReceivers = 100

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	// ...
	dataCh := make(chan int, 100)

	// the sender
	go func() {
		for {
			if value := rand.Intn(MaxRandomNumber); value == 0 {
				// The only sender can close the channel safely.
				close(dataCh)
				return
			} else {
				dataCh <- value
			}
		}
	}()

	// receivers
	for i := 0; i < NumReceivers; i++ {
		tmp := i
		go func() {
			defer wgReceivers.Done()

			// Receive values until dataCh is closed and
			// the value buffer queue of dataCh is empty.
			for value := range dataCh {
				log.Println(value, "\t", tmp)
			}
		}()
	}

	wgReceivers.Wait()
}

// TestChannel1
func TestChannel1(t *testing.T) {
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	readChan(ch)
}

// readChan
func readChan(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

// TestChannel2 使用channel计算100个1相加 无锁化
func TestChannel2(t *testing.T) {
	cur := time.Now()
	ch := make(chan int, 10)
	go func() {
		for i := 0; i < 10000; i++ {
			ch <- 1
		}
		close(ch)
	}()

	var count int
	for v := range ch {
		count += v
	}
	fmt.Println(time.Now().Sub(cur).Microseconds())
	fmt.Println(count)
}
