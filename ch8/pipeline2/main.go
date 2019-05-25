package main

import (
	"fmt"
)

func counter(out chan<- int) {
	for i := 0; i <= 10; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	a := make(chan int)
	b := make(chan int)
	go counter(a)
	go squarer(b, a)
	printer(b)
}
