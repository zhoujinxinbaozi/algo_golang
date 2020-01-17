package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(2) // 设置单核只有一个cpu干活，要么执行完f1，在执行f2，
	                         // 不止一个的时候才会出现并发的场景
	wg.Add(2)
	go f1()
	go f2()
	wg.Wait()
}

func f1() {
	defer wg.Done()
	for i := 1; i < 10; i++ {
		fmt.Printf("f1 ==== %d\n", i)
	}
}

func f2() {
	defer wg.Done()
	for i := 1; i < 10; i++ {
		fmt.Printf("f2 ==== %d\n", i)
	}
}
