// test 类与源文件在同一个包下，文件命名为 文件名_test.go
package testDemo

import (
	"fmt"
	"reflect"
	"testing"
)

// 函数名为 Test开头  参数为 t *testing.T
func TestSplit(t *testing.T) {
	fmt.Println("===== TestSplit =====")
	got := Split("abchdbdd", "b")
	want := []string{"a", "chd", "dd"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:%#v, but got:%#v\n", want, got)
	}
}

// 多个case测试 不需要写多个函数  运行失败可以单独跑一个case
// go test -v 跑所有的测试案例
// go test -v -run=TestGroupSplit/demo1 单独跑一个测试案例
// go test -cover  测试代码覆盖率
// go test -cover -coverprofile=./cover.out  将测试覆盖率存储到文件中
// go tool cover -html=./cover.out 在浏览器中查看哪些代码覆盖或者没覆盖
func TestGroupSplit(t *testing.T) {
	fmt.Println("===== TestGroupSplit =====")
	type testCase struct {
		str  string
		sep  string
		want []string
	}
	testGroup := map[string]testCase{
		"demo1": {"abchdbdd", "b", []string{"a", "chd", "dd"}},
		"demo2": {"abchdbdbd", "b", []string{"a", "chd", "d", "d"}},
	}
	for key, value := range testGroup {
		t.Run(key, func(t *testing.T) {
			got := Split(value.str, value.sep)
			if !reflect.DeepEqual(got, value.want) {
				t.Errorf("want:%#v, but got:%#v\n", got, value.want)
			}
		})
	}

}

// 性能测试
// go test -bench=Split
func BenchmarkSplit(b *testing.B) {
	for i:= 0; i < b.N; i ++{
		Split("abchdbdd", "b")
	}
}
