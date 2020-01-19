package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	for {
		dial, err := net.Dial("tcp", "127.0.0.1:20000")
		if err != nil {
			fmt.Println("dail 127.0.0.1:20000 failed, error:", err)
		}
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("get string failed, error:", err)
		}
		dial.Write([]byte(strings.TrimSpace(line)))
		fmt.Println("send ", strings.TrimSpace(line), " to server")
	}
}
