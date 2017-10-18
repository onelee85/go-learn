package main

import (
	"fmt"
)

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

type IntVect []int

func (v IntVect) Sum() (s int) {
	for _, x := range v {
		s += x
	}
	return
}
func main() {
	r := rect{width: 10, height: 5}
	fmt.Println("area:", r.area())

	rp := &r
	rp.width = 3
	rp.height = 2
	fmt.Println("area:", rp.area())
	fmt.Println("area:", r.area())

	fmt.Println(IntVect{1, 2, 3}.Sum())
}
