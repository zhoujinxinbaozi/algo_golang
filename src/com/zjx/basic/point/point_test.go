/**
 *  @Author: JinxinZhou
 *  @Date  : 2021/2/7 10:59 上午
 */

package point

import (
	"fmt"
	"testing"
)

// TestSlice
func TestSlice(t *testing.T) {
	var slice []int
	for i := 0; i < 100; i++ {
		slice = append(slice, i)
	}
	appendSlice(slice, 101)
	fmt.Println("TestSlice len : ", len(slice))
}

// appendSlice
func appendSlice(slice []int, target int) {
	slice = append(slice, target)
	fmt.Println("appendSlice len : ", len(slice))
}

// TestMap
func TestMap(t *testing.T) {
	var m map[int]int
	addMap(m, 1,1)
	fmt.Println("TestMap len : ", len(m))
}

// addMap
func addMap(m map[int]int, key, value int) {
	if m == nil {
		m = make(map[int]int, 1)
	}
	m[key] = value
	fmt.Println("addMap len : ", len(m))
}
