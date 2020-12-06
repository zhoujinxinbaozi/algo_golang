/**
 *  @Author: JinxinZhou
 *  @Date  : 2020/12/6 11:18 上午
 */

package bytes_convert

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

/*
	二进制字符串转十进制
    十进制转二进制字符串
 */

//0000000000000000000000000000010100000000000000000000000000000000
//0000000000000000000000000000000000000000000000000000000000000010
func TestBytes(t *testing.T) {
	var num int
	num = 17179869184 >> 32 //  小米 4294967297 oppo 17179869184 vivo 21474836480
	bytes := convertToBin(num, 64)
	fmt.Println(bytes)
	c := getInput(bytes)
	out := sq(c)
	sum := 0
	for o := range out {
		sum += o
	}
	fmt.Println(sum)
}

func StringToIntArray(input string) []int {
	output := []int{}
	for _, v := range input {
		output = append(output, int(v))
	}
	for i, j := 0, len(output)-1; i < j; i, j = i+1, j-1 {
		output[i], output[j] = output[j], output[i]
	}
	return output
}

func getInput(input string) <-chan int {
	out := make(chan int)
	go func() {
		for _, b := range StringToIntArray(input) {
			out <- b
		}
		close(out)
	}()

	return out
}
func sq(in <-chan int) <-chan int {
	out := make(chan int)

	var base, i float64 = 2, 0
	go func() {
		for n := range in {
			out <- (n - 48) * int(math.Pow(base, i))
			i++
		}
		close(out)
	}()
	return out
}

// n 十进制的数
// bin 转换成二进制的位数
func convertToBin(n int, bin int) string {
	var b string
	switch {
	case n == 0:
		for i := 0; i < bin; i++ {
			b += "0"
		}
	case n > 0:
		//strcov.Itoa 将 1 转为 "1" , string(1)直接转为assic码
		for ; n > 0; n /= 2 {
			b = strconv.Itoa(n%2) + b
		}
		//加0
		j := bin - len(b)
		for i := 0; i < j; i++ {
			b = "0" + b
		}
	case n < 0:
		n = n * -1
		// fmt.Println("变为整数：",n)
		s := convertToBin(n, bin)
		// fmt.Println("bin:",s)
		//取反
		for i := 0; i < len(s); i++ {
			if s[i:i+1] == "1" {
				b += "0"
			} else {
				b += "1"
			}
		}
		// fmt.Println("~bin :",b)
		//转化为整形，之后加1 这里必须要64，否则在转化过程中可能会超出范围
		n, err := strconv.ParseInt(b, 2, 64)
		if err != nil {
			fmt.Println(err)
		}
		//转为bin
		//+1
		b = convertToBin(int(n+1), bin)
	}
	return b
}
