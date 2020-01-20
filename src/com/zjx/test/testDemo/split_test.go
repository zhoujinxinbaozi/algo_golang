// test 类与源文件在同一个包下，文件命名为 文件名_test.go
package testDemo

import (
	"reflect"
	"testing"
)
// 函数名为 Test开头  参数为 t *testing.T
func TestSplit(t *testing.T) {
	got := Split("abchdbdd", "b")
	want := []string{"a", "chd", "dd"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:%#v, but got:%#v\n", want, got)
	}
}
