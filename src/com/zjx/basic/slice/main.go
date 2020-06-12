package main

import (
	"fmt"
)

/**
 * @Auther: zhoujinxin@bytedance.com
 * @Date: 2020/6/11 11:50
 */

func main() {

	b := []int64{1, 23}
	fmt.Println("cap of b before ", cap(b)) // 2
	b = append(b, 4, 5, 6)
	fmt.Println("cap of b after ", cap(b)) // 6

	c := []int32{1, 23}
	fmt.Println("cap of c before ", cap(c)) //2
	c = append(c, 4, 5, 6)
	fmt.Println("cap of c after ", cap(c)) // 8

	//cc := []int64{1, 2}
	//for i := 3; i < 10; i++ {
	//	fmt.Print(len(cc))
	//	cc = append(cc, int64(i))
	//	fmt.Println("\t", cap(cc))
	//}

	// 协程不安全golang编译器可以检查   不用mutex 不可以执行
	//var m map[int]int
	//m = make(map[int]int)
	//var wg sync.WaitGroup
	//var mutex sync.Mutex

	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	for i := 0; i < 100; i++ {
	//		mutex.Lock()
	//		m[i] = i
	//		mutex.Unlock()
	//	}
	//}()
	//
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	for i := 0; i < 100; i++ {
	//		mutex.Lock()
	//		m[i+200] = i + 200
	//		mutex.Unlock()
	//	}
	//}()
	//
	//wg.Wait()
	//fmt.Println(len(m))

	//Test1()
}

// 测试slice扩容和内存对齐
func Test1() {
	slice := []int64{1, 2}
	for i := 3; i < 60; i++ {
		fmt.Print(len(slice))
		slice = append(slice, int64(i))
		fmt.Println("\t", cap(slice))
	}
}
