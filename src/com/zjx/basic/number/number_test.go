/**
 *  @Author: JinxinZhou
 *  @Date  : 2021/5/17 11:07 上午
 */

package number

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/imroc/biu"
	"github.com/stretchr/testify/assert"
)

// TestUint64Sub uint64类型 相减
func TestUint64Sub(t *testing.T) {
	var s1 uint64 = 5
	var s2 uint64 = 1000
	var s3 uint64 = 5600
	fmt.Println(s1 > (s2 - s3))
	i := s2 - s3
	fmt.Println(i) // 18446744073709547016
}

func TestBit(t *testing.T) {
	var result uint64
	result += 1 << 1
	assert.Equal(t, uint64(2), result)

	result = result << 1 + 1
	assert.Equal(t, uint64(5), result)

	result = 0
	result = 1 << 2 + 1 << 1 + 1
	assert.Equal(t, uint64(7), result)
}

func TestInt2Bytes(t *testing.T) {
	fmt.Println(biu.ToBinaryString(int32(5)))
}

func TestEncodeBit(t *testing.T)  {
	//var num uint32 = 5
	//fmt.Println(unsafe.Sizeof(num))
	//fmt.Println(biu.ToBinaryString(num))
	//fmt.Println(biu.ToBinaryString(EncodeBit(num, 3, 3)))
	infoBit := uint32(0)
	infoBit |= uint32(1)
	fmt.Println(infoBit)
}

func EncodeBit(num, offset, len uint32) uint32 {
	shiftLen := 8*uint32(unsafe.Sizeof(num)) - len
	num = (num << shiftLen) >> shiftLen // 截取长度为len的bit值
	num = num << offset
	return num
}

func TestMap(t *testing.T) {
	m := make(map[int]int)
	m[0] = 1
	v, ok := m[0]
	fmt.Println("ok : ", ok, " v : ", v)
}

func TestChannel(t *testing.T) {
	var c chan<- []string
	for {
		select {
		case c <- []string{"1"}:
			//fmt.Println("write success")
		default:
			fmt.Println("error")
		}
	}
}
