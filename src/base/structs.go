package main

import (
	"fmt"
	"strings"
)

type person struct {
	name string
	age  int
}

func upName(p *person) {
	p.name = strings.ToUpper(p.name)
}

func main() {
	//struct as a pointer
	p1 := new(person)
	p1.age = 20
	fmt.Println(p1)

	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(&person{name: "Ann", age: 40}) //struct as a literal

	p := person{name: "Sean", age: 50}
	fmt.Println(p.name)

	sp := &p

	fmt.Println(sp.age)

	sp.age = 40
	fmt.Println(p.age)

	upName(sp)
	fmt.Println(p.name)
	fmt.Println(sp.name)

	p2 := person{"James", 30}
	upName(&p2)
	fmt.Println(p2)
}
