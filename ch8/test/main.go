package main

import (
	"fmt"
	// "time"
	// "time"
)

// 先准备好recv,再send,不然send会一直阻塞
func main() {
	ch := make(chan int)

	go func() {
		fmt.Println("11111")
		ch <- 11111
		// close(ch)
	}()

	v := <-ch
	// vv := <-ch
	// b := <-ch

	fmt.Println("2222", v)
	// fmt.Println("2222", v, vv, b)
}
