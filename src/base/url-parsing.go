package main

import (
	"fmt"
	"net/url"
)

var p = fmt.Println

func main() {
	s := "https://user.2323/path?q=123#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	p(u.Scheme)
	p(u.Hostname())
	p(u.RequestURI())
	p(u.RawQuery)
	m, e := url.ParseQuery(u.RawQuery)
	p(e)
	p(m["q"])
}
