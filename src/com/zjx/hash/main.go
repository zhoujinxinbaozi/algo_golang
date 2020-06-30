package main

import (
	"fmt"
	"hash/crc32"
)

/**
 * @Auther: zhoujinxin@bytedance.com
 * @Date: 2020/6/19 11:16
 */

// String hashes a string to a unique hashcode.
//
// crc32 returns a uint32, but for our use we need
// and non negative integer. Here we cast to an integer
// and invert it if the result is negative.
func String(s string) int {
	v := int(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	// v == MinInt
	return 0
}

func main() {
	str := "1;1,3,5;2;3,5,7"
	str1 := "1;1,3,5;3;3,5,7"

	hc := String(str)
	fmt.Println("hashcode:", hc)

	hc1 := String(str1)
	fmt.Println("hashcode:", hc1)

}
