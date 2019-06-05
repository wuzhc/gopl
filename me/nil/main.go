package main

import (
	"fmt"
)

func main() {
	// cannot convert nil to type struct {}
	// var v1 struct{}
	// if v1 == nil {
	// 	fmt.Println("struct is nil")
	// }

	// cannot convert nil to type bool
	// var v7 bool
	// if v7 == nil {
	// fmt.Println("bool is nil")
	// }

	// v2 interface{} == nil
	var v2 interface{}
	if v2 == nil {
		fmt.Println("v2 interface{} == nil")
	}

	// var v3 []int == nil
	var v3 []int
	if v3 == nil {
		fmt.Println("var v3 []int == nil")
	}

	// var v4 = make([]int,0) != nil
	var v4 = make([]int, 0)
	if v4 == nil {
		fmt.Println("var v4 = make([]int,0) == nil")
	} else {
		fmt.Println("var v4 = make([]int,0) != nil")
	}

	// var v5 map[string]int == nil
	var v5 map[string]int
	if v5 == nil {
		fmt.Println("var v5 map[string]int == nil")
	}

	// var v6 = make(map[string]int) != nil
	var v6 = make(map[string]int)
	if v6 == nil {
		fmt.Println("var v6 = make(map[string]int) == nil")
	} else {
		fmt.Println("var v6 = make(map[string]int) != nil")
	}

	// var v8 chan == nil
	var v8 chan int
	if v8 == nil {
		fmt.Println("var v8 chan == nil")
	} else {
		fmt.Println("var v8 chan != nil")
	}

	// var v9 = make(chan int) != nil
	var v9 = make(chan int)
	if v9 == nil {
		fmt.Println("var v9 = make(chan int) == nil")
	} else {
		fmt.Println("var v9 = make(chan int) != nil")
	}

	// var v10 func() == nil
	var v10 func()
	if v10 == nil {
		fmt.Println("var v10 func() == nil")
	} else {
		fmt.Println("var v10 func() != nil")
	}

	// var v11 = func() {} == nil
	var v11 = func() {}
	if v11 == nil {
		fmt.Println("var v11 = func() {} == nil")
	} else {
		fmt.Println("var v11 = func() {} == nil")
	}

	// 用append的结果是[0 0 1 2 3]
	var v12 = make([]int, 2)
	v12 = append(v12, 1, 2, 3)
	fmt.Println(v12)
}
