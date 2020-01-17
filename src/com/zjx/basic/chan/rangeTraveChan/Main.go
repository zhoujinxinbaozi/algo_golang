package main

import "fmt"

func main() {
	f1()
}

func f2(){
	var c1 chan int
	c1 = make(chan int)
	go func() {
		for i := 1; i < 100; i++ {
			c1 <- i
		}
		close(c1) // 必须关闭channel，否则会有问题
	}()

	for c := range c1 {
		fmt.Println(c)
	}
}

func f1(){
	var c1 chan int
	c1 = make(chan int)
	go func() {
		for i := 1; i < 100; i++ {
			c1 <- i
		}
		close(c1)
	}()
	for{
		if x, ok := <-c1; ok{ // chan关闭ok为false
			fmt.Println(x)
		}else{
			break
		}
	}
}
