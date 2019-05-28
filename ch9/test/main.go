package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		for c := range ch {
			fmt.Println(c)
		}
	}()

	// for i := 0; i < 10; i++ {
	// ch <- i
	// }
	// close(ch)
	time.Sleep(3 * time.Second)
}
