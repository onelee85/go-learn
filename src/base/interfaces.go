package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	permi() float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 { //必须实现接口所有方法
	return r.width * r.height
}
func (r rect) permi() float64 { //必须实现接口所有方法
	return 2*r.width + 2*r.height
}

type circle struct {
	radius float64
}

func (c *circle) area() float64 { //指针类型
	return math.Pi * c.radius * c.radius
}

func (c circle) permi() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.permi())
}

func main() {
	r := rect{width: 10, height: 5}
	c := circle{radius: 5}
	measure(r)
	var c1 = &c
	c1.area()
	c.permi()
	measure(&c)
}
