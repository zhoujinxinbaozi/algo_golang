/**
 *  @Author: JinxinZhou
 *  @Date  : 2020/10/28 2:47 下午
 */

package _map

import (
	"fmt"
	"testing"
	"time"
)

var m map[string]string
func TestMap(t *testing.T) {
	m = make(map[string]string)
	var key, value string
	key = "zhoujinxin"
	value = "hhh"

	go func() {
		for {
			write("lll", value)
		}
	}()

	go func() {
		for {
			write(key, value)
			//fmt.Println(value)
		}
	}()

	time.Sleep(time.Duration(1) * time.Second)
}

func read(name string) {
	fmt.Println(m[name])
}

func write(name, value string) {
	m[name] = value
}
