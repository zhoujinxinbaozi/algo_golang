/**
 *  @Author: JinxinZhou
 *  @Date  : 2020/12/7 3:53 下午
 */

package _defer

import (
	"fmt"
	"os"
	"testing"
)

// TestDefer
func TestDefer(t *testing.T) {

	fd, err := os.Create("./ss/tmp")
	if err != nil {
		fmt.Println("err")
		fmt.Println(fd == nil)
		return
	}

	defer func() {
		fmt.Println("lalalal")
		fd.Close()
	}()
}


