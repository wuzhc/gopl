package main

import (
	"fmt"
	"time"
)

var an, bn int

func main() {
	a := make(chan struct{})
	b := make(chan struct{})

	go func() {
		for {
			select {
			case <-b:
				bn++
			default:
				a <- struct{}{}
			}
		}
	}()

	go func() {
		for {
			select {
			case <-a:
				an++
			default:
				b <- struct{}{}
			}
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Printf("an:%v, bn:%v\n", an, bn)
}
