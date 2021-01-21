package main

import (
	"fmt"
	"time"

	"com/zjx/algo_golang/src/com/zjx/basic/log/logEntity"
)

func main() {
	for {
		logger, err := logEntity.NewLogger("debug")
		if err != nil {
			fmt.Printf("inital logEntity failed, error:%v", err)
			return
		}
		logger.Debug("this is debug log")
		logger.Info("this is info log")
		logger.Warning("this is warning log")
		logger.Error("this is error log")
		time.Sleep(time.Second * 1)
	}

}
