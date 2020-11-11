/**
 *  @Author: JinxinZhou
 *  @Date  : 2020/11/11 2:43 下午
 */

package _interface

import (
	"fmt"
	"testing"
)

func TestInterfaceJudge(t *testing.T) {

}

func receiveAnyThing(i interface{}) {
	switch i.(type) {
	case []int:
		fmt.Println("type is []int, value : ", i.([]int))
	case string:
		fmt.Println("type is string, value : ", i.(string))
	}
}
