package main

import (
	"fmt"
	"strconv"
	"testing"
)

/**
 * @Auther: zhoujinxin@bytedance.com
 * @Date: 2020/6/12 16:59
 */

func TestSaveN(t *testing.T) {
	SaveN(6.8956, 2)
}

// 对传进来的小数，保留n位小数
func SaveN(target float64, n int64) {
	n2, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", target), 64)
	fmt.Println(fmt.Sprintf("%.2f", target))
	fmt.Println(n2)
}
