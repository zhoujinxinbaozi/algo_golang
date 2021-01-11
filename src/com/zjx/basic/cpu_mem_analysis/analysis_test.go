/**
 *  @Author: JinxinZhou
 *  @Date  : 2020/12/16 3:03 下午
 */

package cpu_mem_analysis

import (
	"fmt"
	"os"
	"runtime/pprof"
	"testing"
)

func TestCpuAnalysis(t *testing.T) {
	f, err := os.Create("./cpuAnalysis")
	if err != nil {
		fmt.Printf("err : %+v\n", err)
		return
	}

	// cpu
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	// memory
	pprof.WriteHeapProfile(f)

	deadCircle()
	deadCircle1()
}

func deadCircle() {
	for i := 0; i < 10000; i ++ {
		for j := 0; j < 10000; j ++ {

		}
	}
}

func deadCircle1() {
	for j := 0; j < 1000000; j ++ {
		for i := 0; i < 10000; i ++ {

		}
	}
}

func TestMap(t *testing.T) {
	m := make(map[int]int)
	if v, ok := m[1]; ok {
		m[1] = v + 1
	} else {
		m[1]++
		fmt.Println("v : ", v)
	}
	fmt.Println(m)
}

func TestSlice(t *testing.T) {
	var s []int
	s = append(s, 1)
	s = append(s, 2)
	s = append(s, 3)
	fmt.Println(s[:len(s)])
}