/**
 *  @Author: JinxinZhou
 *  @Date  : 2022/1/26 2:32 下午
 */

package poll

import (
	"fmt"
	"sync"
	"testing"
)

// 对象的复用 池子里有的话 直接复用 没有的话 创建新的结构体
var pool *sync.Pool

type Person struct {
	Name string
}

func initPool() {
	pool = &sync.Pool {
		New: func()interface{} {
			fmt.Println("Creating a new Person")
			return new(Person)
		},
	}
}

func TestPoll(t *testing.T) {
	initPool()

	p := pool.Get().(*Person)
	fmt.Println("首次从 pool 里获取：", p)

	p.Name = "first"
	fmt.Printf("设置 p.Name = %s\n", p.Name)

	pool.Put(p)

	fmt.Println("Pool 里已有一个对象：&{first}，调用 Get: ", pool.Get().(*Person))
	fmt.Println("Pool 没有对象了，调用 Get: ", pool.Get().(*Person))
}
