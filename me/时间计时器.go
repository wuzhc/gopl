package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second * 7)
			c <- false
		}

		time.Sleep(time.Second * 7)
		c <- true
	}()

	go func() {
		timer := time.NewTimer(time.Second * 5)

		for {
			// try to read from channel, block at most 5s.
			// if timeout, print time event and go on loop.
			// if read a message which is not the type we want(we want true, not false),
			// retry to read.

			if !timer.Stop() {
				fmt.Println("timer stop")
				select {
				case <-timer.C:
				default:
				}
			} else {
				fmt.Println("timer no stop")
			}

			if timer.Reset(time.Second) {
				fmt.Println("如果调用时timer还在等待中会返回真")
			} else {
				fmt.Println("如果timter已经到期或者被停止了会返回假")
			}

			select {
			case b := <-c:
				if b == false {
					fmt.Println(time.Now(), ":recv false. continue")
					continue
				}
				//we want true, not false
				fmt.Println(time.Now(), ":recv true. return")
				return
			case <-timer.C:
				fmt.Println(time.Now(), ":timer expired")
				continue
			}
		}
	}()

	//to avoid that all goroutine blocks.
	var s string
	fmt.Scanln(&s)
}
