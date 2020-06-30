package copy

import (
	"fmt"
	"testing"
)

/**
 * @Auther: zhoujinxin@bytedance.com
 * @Date: 2020/6/26 16:40
 */

func TestCopy(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := make([]int, len(s1))
	copy(s2, s1)
	s1[0] = 9
	s2[0] = 10
	fmt.Printf("%p\n", &s1)
	fmt.Printf("%p\n", &s2)
	fmt.Printf("%v\n", s1)
	fmt.Printf("%v\n", s2)
}

func TestStr(t *testing.T) {
	s := "123gh"
	for _, value := range s {
		fmt.Println(string(value))
	}
}
