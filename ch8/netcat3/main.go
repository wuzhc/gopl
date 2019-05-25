package main

import (
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

	done := make(chan int)
	go func() {
		if n, err := io.Copy(os.Stdout, c); err != nil {
			log.Fatalln(err)
		} else {
			log.Println("read ", n)
		}
		done <- 1
	}()

	mustCopy(c, os.Stdin)
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if n, err := io.Copy(dst, src); err != nil {
		log.Fatalln(err)
	} else {
		log.Println("read ", n)
	}
}
