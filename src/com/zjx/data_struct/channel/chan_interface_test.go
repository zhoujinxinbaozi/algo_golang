/**
 *  @Author: JinxinZhou
 *  @Date  : 2022/1/27 12:22 下午
 */

package channel

import (
	"fmt"
	"testing"
	"time"
)

type student struct {
	Name string
}

type teacher struct {
	Name string
}

// TestChanInterface 测试interface类型的转换通过channel
func TestChanInterface(t *testing.T) {
	ch := make(chan interface{}, 5)
	go func() {
		ch <- &student{Name: "zhoujinxin"}
		ch <- &teacher{Name: "zjx"}
		time.Sleep(2 * time.Second)
		close(ch)
	}()

	for ss := range ch {
		if student, ok := ss.(*student); ok {
			fmt.Printf("student name : %v\n", student.Name)
		} else if teacher, ok := ss.(*teacher); ok {
			fmt.Printf("teacher name : %v\n", teacher.Name)
		} else {
			fmt.Println("other")
		}
	}
}
