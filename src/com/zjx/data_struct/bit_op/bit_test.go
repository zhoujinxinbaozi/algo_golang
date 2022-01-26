/**
 *  @Author: JinxinZhou
 *  @Date  : 2022/1/24 5:23 下午
 */

package bit_op

import (
	"fmt"
	"testing"
	"unsafe"
)

const (
	BIT_LEN_OF_BYTE = 8

	BIT_OFFSET8  uint32 = 8

	VERSION_FIRST  uint32 = 24
	VERSION_SECOND uint32 = 16
	VERSION_THIRD  uint32 = 8
	VERSION_FOURTH uint32 = 0
)

var versionMap = map[int]uint32{
	0: VERSION_FIRST,
	1: VERSION_SECOND,
	2: VERSION_THIRD,
	3: VERSION_FOURTH,
}

func TestBit(t *testing.T) {
	versionList := []uint32{127, 0, 0, 1}
	var result uint32
	for index, version := range versionList {
		result |= EncodeBit(version, versionMap[index], BIT_OFFSET8)
	}
	fmt.Println("result : ", result)

	fmt.Println("version bit first", DecodeBit(uint64(result), uint64(VERSION_FIRST), uint64(BIT_OFFSET8)))
	fmt.Println("version bit second", DecodeBit(uint64(result), uint64(VERSION_SECOND), uint64(BIT_OFFSET8)))
	fmt.Println("version bit third", DecodeBit(uint64(result), uint64(VERSION_THIRD), uint64(BIT_OFFSET8)))
	fmt.Println("version bit fourth", DecodeBit(uint64(result), uint64(VERSION_FOURTH), uint64(BIT_OFFSET8)))
}

// DecodeBit 反解bit值
func DecodeBit(num, offset, len uint64) uint64 {
	shiftLen := BIT_LEN_OF_BYTE*uint64(unsafe.Sizeof(num)) - len
	num = num >> offset
	num = (num << shiftLen) >> shiftLen // 截取长度为len的bit值
	return num
}

// EncodeBit
func EncodeBit(num, offset, len uint32) uint32 {
	shiftLen := BIT_LEN_OF_BYTE*uint32(unsafe.Sizeof(num)) - len
	num = (num << shiftLen) >> shiftLen // 截取长度为len的bit值
	num = num << offset
	return num
}
