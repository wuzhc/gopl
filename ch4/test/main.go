package main

import (
	"fmt"
)

func GetName() string {
	defer func() {
		fmt.Println("hello ")
	}()
	fmt.Println("good")
	return "wuzhc"
}

func main() {
	name := GetName()
	fmt.Println(name)
}
