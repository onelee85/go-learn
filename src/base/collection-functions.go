package main

import (
	"fmt"
	"strings"
)

var p = fmt.Println

func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func main() {
	var strs = []string{"peach", "apple", "pear", "plum"}
	p(Index(strs, "pear"))
	p(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))
}
