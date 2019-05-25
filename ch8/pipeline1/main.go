package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	d := make(chan int)

	go func() {
		for i := 0; i <= 10; i++ {
			c <- i
		}
		close(c)
	}()

	go func() {
		for {
			x, ok := <-c
			if !ok {
				break
			}
			d <- x * x
		}
		close(d)
	}()

	for {
		v, ok := <-d
		if !ok {
			break
		}
		fmt.Println("结果是:", v)
	}

	fmt.Println("end.......")
}
