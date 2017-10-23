package main

import (
	"fmt"
	"regexp"
)

var p = fmt.Println

func main() {
	r, _ := regexp.Compile("p([a-z]+)ch")

	p(r.MatchString("peach"))
	p(r.FindString("peach punch"))
	p(r.FindAllString("peach punch pinch", -1))
}
