package main

import (
	"encoding/json"
	"fmt"
)

var p = fmt.Println

type Response struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	bolB, _ := json.Marshal(true)
	p(string(bolB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	p(string(mapB))

	res := &Response{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}

	resJson, _ := json.Marshal(res)
	p(string(resJson))

	str := `{"page":1, "fruits":["apple", "peach"]}`
	res1 := Response{}
	json.Unmarshal([]byte(str), &res1)
	p(res1)
	p(res1.Fruits[0])
}
