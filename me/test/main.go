package main

import (
	"fmt"
	"time"
)

func main() {
	var i = 1
	time.Sleep(time.Duration(i) * time.Second)
	fmt.Println("timeout")
}
