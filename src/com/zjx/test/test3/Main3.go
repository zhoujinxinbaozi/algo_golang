package main

import "fmt"

/**
 * @Auther: zhoujinxin@bytedance.com
 * @Date: 2020/5/29 16:37
 */
// chan 可以认为就是一个函数，是值传递
func main() {
	ch := make(chan int, 4)
	i := 1
	ch <- i
	j := <-ch
	fmt.Println(&i)
	fmt.Println(&j)
}
