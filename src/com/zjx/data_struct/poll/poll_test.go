/**
 *  @Author: JinxinZhou
 *  @Date  : 2022/1/26 2:32 下午
 */

package poll

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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

	p := pool.Get().(*Person) // 池子里没有对象 创建新的对象
	assert.Equal(t, "", p.Name)

	p.Name = "zhoujinxin"
	pool.Put(p) // 池子里有一个对象

	p = pool.Get().(*Person) // 取出来直接用
	assert.Equal(t, "zhoujinxin", p.Name)

	p = pool.Get().(*Person) // 池子里没有对象 创建新的对象
	assert.Equal(t, "", p.Name)
}
