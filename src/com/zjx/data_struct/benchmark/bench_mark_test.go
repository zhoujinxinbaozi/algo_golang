/**
 *  @Author: JinxinZhou
 *  @Date  : 2022/1/27 12:56 下午
 */

package benchmark

import (
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