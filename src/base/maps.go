package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map :", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	v, exists := m["k2"]
	fmt.Println("exists:", v, exists)

	m2 := map[string]int{"a": 1, "b": 2}
	fmt.Println("m2:", m2)
}
