package generics

import (
	"fmt"
	"testing"
)

// 定义一个int类型的集合类型 GenericsInt
type GenericsInt interface {
	int | int32 | int64
}

func TestGenericsInt(t *testing.T) {
	fmt.Println(IntAddOne(int32(1)))
	fmt.Println(IntAddOne(int64(1)))

	fmt.Println(MaxInt(1, 100))
}

// 对于GenericsInt类型中的变量进行+1操作
func IntAddOne[T GenericsInt](target T) T {
	return target + 1
}

func MaxInt[T GenericsInt](m, n T) T {
	if m < n {
		return n
	}
	return m
}
