package main

import (
	"fmt"
)

type Point struct {
	x, y int
}

type Circle struct {
	Point
	radius int
}

type Wheel struct {
	Circle
	spokes int
}

func main() {
	var w Wheel
	w.x = 8
	w.y = 8
	w.radius = 10
	w.spokes = 20

	v := Wheel{Circle{Point{1, 1}, 1}, 1}

	fmt.Printf("%v----%v\n", w, v)
}
