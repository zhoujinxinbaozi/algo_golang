package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen 127.0.0.1:20000 failed, error:", err)
	}

	for {
		accept, err := listen.Accept()
		fmt.Println("wait accept ...")
		if err != nil {
			fmt.Println("accept 127.0.0.1:20000 failed, error:", err)
		}
		var byteBuffer [128]byte
		n, err := accept.Read(byteBuffer[:])
		if err != nil {
			fmt.Println("read 127.0.0.1:20000 failed, error:", err)
		}
		go process(byteBuffer, n)
	}

}

func process(byteBuffer [128]byte, n int) {
	fmt.Println("get data : ", string(byteBuffer[:n]))
}
