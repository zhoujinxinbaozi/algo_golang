/**
 *  @Author: JinxinZhou
 *  @Date  : 2022/1/27 12:56 下午
 */

package benchmark

import (
	"fmt"
	"testing"
)

// go test -v ./src/com/zjx/data_struct/benchmark -bench=.

// BenchmarkRepeat Benchmark 开头
// param b *testing.B
// param b.N
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat(`b`, 5)
	}
}

func Repeat(char string, count int) (result string) {
	for i := 0; i < count; i++ {
		result += char
	}
	return
}

type CustomError struct {
	Code int64
	Msg  string
}

func (p *CustomError) Error() string {
	return fmt.Sprintf("code : %v, msg : %v", p.Code, p.Msg)
}

func TestError(t *testing.T) {
	err := execute()
	fmt.Println(err == nil) // true or false

	err = execute1()
	fmt.Println(err == nil) // true or false

	err = execute2()
	fmt.Println(err == nil) // true or false
}

func execute() error {
	var err *CustomError
	// do something
	return err
}

func execute1() error {
	var err *CustomError
	// do something
	if err != nil {
		return err
	}
	return nil
}

func execute2() *CustomError {
	var err *CustomError
	// do something

	return err
}

func TestInterfaceAssign(t *testing.T) {
	var i interface{}
	userID := int64(1)
	i = userID
	fmt.Println(i)
}
