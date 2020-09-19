package single

import (
	"fmt"
	"sync"
	"testing"
)

// TestSingle 单例测试
func TestSingle(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i ++ {
		go func() {
			defer wg.Done()
			t := GetSingleInstance()
			fmt.Printf("%p\n", t)
		}()
	}
	wg.Wait()
}

type Instance struct {

}


var instance *Instance

// GetSingleInstance 获取单例
func GetSingleInstance() *Instance {
	var mutex sync.Mutex
	if instance == nil {
		mutex.Lock()
		if instance == nil {
			instance = &Instance{}
		}
	}
	return instance
}
