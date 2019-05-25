package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	c, err := net.Dial("tcp", "127.0.0.1:9502")
	if err != nil {
		log.Fatalln(err)
	} else {
		defer c.Close()
	}

	go mustCopy(os.Stdout, c)
	mustCopy(c, os.Stdin)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if n, err := io.Copy(dst, src); err != nil {
		log.Fatalln(err)
	} else {
		log.Println("read ", n)
	}
}
