package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	c, err := net.Dial("tcp", "127.0.0.1:9501")
	if err != nil {
		log.Fatalln(err)
	} else {
		defer c.Close()
	}

	// 读取服务端数据,输出到标准输出
	// n, err := io.Copy(os.Stdout, c)
	// if err != nil {
	// 	log.Fatalln(err)
	// } else {
	// 	log.Println("read ", n)
	// }

	mustCopy(os.Stdout, c)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if n, err := io.Copy(dst, src); err != nil {
		log.Fatalln(err)
	} else {
		log.Println("read ", n)
	}
}
