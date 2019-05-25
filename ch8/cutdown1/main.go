package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var n = 0
	abort := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1))
		fmt.Println("read read", n)
		abort <- struct{}{}
	}()

	tick := time.NewTicker(1 * time.Second)
	for {
		n++
		select {
		case <-tick.C:
			fmt.Println(n)
		case <-abort:
			fmt.Println("all done")
			return
		}
	}
}
