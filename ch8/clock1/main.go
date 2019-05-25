package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:9501")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for {
		c, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		go func() {
			defer c.Close()
			for {
				n, err := io.WriteString(c, time.Now().Format("2006-01-02 15:04:05")+"\n")
				if err != nil {
					return
				}

				time.Sleep(1 * time.Second)
				fmt.Println("write ", n)
			}
		}()
	}
}
