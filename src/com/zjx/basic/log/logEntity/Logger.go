package logEntity

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
	"time"
)

type logLevel int8

const (
	DEBUG   string   = "DEBUG"
	INFO    string   = "INFO"
	WARNING string   = "WARNING"
	ERROR   string   = "ERROR"
	debug   logLevel = iota
	info
	warning
	err
)

var fileObj *os.File

/*
	根据传进来的字符串进行映射
*/
func getLevel(level string) logLevel {
	level = strings.ToUpper(level)
	switch level {
	case DEBUG:
		return debug
	case INFO:
		return info
	case WARNING:
		return warning
	case ERROR:
		return err
	default:
		return logLevel(-1)
	}
}

/*
	日志对象
*/
type logger struct {
	level logLevel
}

/*
	初始化日志对象
*/
func NewLogger(level string) (logger, error) {
	// 初始化文件句柄
	filePath := "/Users/apple/workspace/coding/src/com/zjx/basic/log"
	file, err2 := os.OpenFile(path.Join(filePath, "bytedance.log"), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err2 != nil {
		str := fmt.Sprintf("open file failed, error:%v\n", err2)
		return logger{level: -1}, errors.New(str)
	}
	fileObj = file
	levelInt := getLevel(level)
	if levelInt == -1 {
		str := fmt.Sprint("the type is not correct", level)
		return logger{level: -1}, errors.New(str)
	}
	return logger{level: levelInt}, nil
}

func (l logger) Debug(msg string) {
	if debug >= l.level {
		print(DEBUG, msg)
	}
}

func (l logger) Info(msg string) {
	if info >= l.level {
		print(INFO, msg)
	}
}

func (l logger) Warning(msg string) {
	if warning >= l.level {
		print(WARNING, msg)
	}
}

func (l logger) Error(msg string) {
	if err >= l.level {
		print(ERROR, msg)
	}
}

/*
	日志打印
*/
func print(level string, msg string) {
	now := time.Now()
	// 向console中打印
	str := fmt.Sprintf("[%s] [%s] [%s]\n", level, now.Format("2006-01-02 15:04:05"), msg)
	fmt.Printf(str)
	fileObj.WriteString(str)
}
