package main

import (
	"errors"
	"fmt"
	"time"
	// "time"
)

// 通道用于不同goroutine之间通信,同个goroutine不能使用通道
func main() {
	ticker := time.NewTicker(5 * time.Second)
	select {
	case s, err := test():
		fmt.Println(s)
	case <-ticker.C:
		fmt.Println("timeout")
	}
}

func test() (string, error) {
	time.Sleep(time.Second * 1)
	return "wuzhc", errors.New("nothing")
}
