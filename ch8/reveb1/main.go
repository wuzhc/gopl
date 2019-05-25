package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:9502")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		c, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handleConn(c)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()

	input := bufio.NewScanner(c)

	ch := make(chan struct{})
	go func() {
		for input.Scan() {
			ch <- struct{}{}
		}
	}()

	for {
		select {
		case <-ch:
			go echo(c, input.Text(), 2*time.Second)
		case <-time.After(10 * time.Second):
			// 10秒后客户端没有操作,则断开链接
			fmt.Println("close...")
			return
		}
	}
}

func echo(c net.Conn, input string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToLower(input))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", input)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToUpper(input))
}
