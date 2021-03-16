package fsnotify

import (
	"fmt"
	"testing"

	"github.com/fsnotify/fsnotify"
)

// TestFsNotify fsnotify 测试
func TestFsNotify(t *testing.T) {
	// new watcher object
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Printf("[ fsnotify.NewWatcher failed, err : %+v\n ]", err)
		return
	}
	defer watcher.Close()

	// add monitor object
	err = watcher.Add("/Users/zhoujinxin/go/src/github/algo_golang/src/com/zjx/basic/fsnotify/fsnotify.txt")
	if err != nil {
		fmt.Printf("[ watcher.Add failed, err : %+v\n ]", err)
		return
	}

	// 监控对象的变化
	go MonitorObject(watcher)

	var ch chan string
	ch = make(chan string)
	<-ch

}

// MonitorObject 监控对象
// @param watcher 监控对象
// @return nil
func MonitorObject(watcher *fsnotify.Watcher) {
	for {
		select {
		case ev := <-watcher.Events:
			if ev.Op&fsnotify.Write != fsnotify.Write {
				fmt.Printf("not write event, op : %+v, name : %+v\n", ev.Op, ev.Name)
			} else {
				fmt.Printf("write event, op : %+v, name : %+v\n", ev.Op, ev.Name)
			}
		case err := <-watcher.Errors:
			fmt.Printf("watcher has err : %+v\n", err)
			return
		}
	}
}

type Op int32

const (
	W = iota + 1
	WW
	WWW
)

// TestIota
func TestIota(t *testing.T) {
	fmt.Printf("%T, %v\n", W, W)
	fmt.Println(WW)
	fmt.Println(WWW)
}
