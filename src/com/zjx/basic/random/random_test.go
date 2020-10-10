/**
 *  @Author: JinxinZhou
 *  @Date  : 2020/10/9 2:31 下午
 */

package random

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRandom(t *testing.T){
	rand.Seed(time.Now().Unix())
	fmt.Println("My first lucky number is", rand.Intn(10))
	fmt.Println("My second lucky number is", rand.Intn(10))
	fmt.Println(time.Now().Unix()-86400)
}
