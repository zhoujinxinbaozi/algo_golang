package main

import (
	"flag"
	"fmt"
)

// ./flag -name=zhoujixin
// 获取命令行指定名字的参数
func main() {
	var name string
	flag.StringVar(&name, "name", "", "请输入姓名")
	flag.Parse() // 必须写
	fmt.Println(name)
}
