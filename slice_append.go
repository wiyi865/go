package main

import "fmt"

type p map[string][]string
var b []string
func main() {
	a := p{ "name": {"xiejun"},
		"skill": {
			"数学",
			"语文",
			"计算机",
		},
		"age": {"15"},
	}

	for _, v := range a{
		for _, m := range v{
		b = append(b, m)
		}
	}
	fmt.Println(b)
}
