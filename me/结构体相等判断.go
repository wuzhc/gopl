package main

import (
	"fmt"
)

type Stu struct {
	name string
}

func NewStu(name string) Stu {
	return Stu{name}
}

func NewStuPtr(name string) *Stu {
	return &Stu{name}
}

func main() {
	stu1 := NewStu("wuzhc")
	stu2 := NewStu("wuzhc")
	fmt.Println(stu1 == stu2) // true

	stu3 := NewStuPtr("wuzhc")
	stu4 := NewStuPtr("wuzhc")
	fmt.Println(stu3 == stu4) // false
}
