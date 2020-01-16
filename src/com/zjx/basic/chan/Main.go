package main

import (
	"fmt"
	"sync"
)

var c1 chan int
var c2 chan int
var wg sync.WaitGroup

/*
	创建10个数发送到c1中，在从c1中取出，并求平方，并存入到c2中，进行打印
 */
func main() {
	c1 = make(chan int, 5)
	c2 = make(chan int, 5)
	wg.Add(3)
	go createNumber()
	go pingfa()
	go print()
	wg.Wait()
}

func createNumber() {
	defer wg.Done()
	for i := 1; i < 10; i++ {
		c1 <- i
	}
	fmt.Println("close c1")
	close(c1)// 不关闭会有问题
}

func pingfa() {
	defer wg.Done()
	for i := range c1{
		c2 <- i * i
	}
	fmt.Println("close c2")
	close(c2)// 不关闭会有问题
}

func print() {
	defer wg.Done()
	for i := range c2 {
		fmt.Printf("the value is %d\n", i)
	}
	fmt.Println("end")
}
