package main

import (
	"fmt"
)

func main() {
	m := make(map[string]string)
	m["name"] = "xiaowang"
	fmt.Println(m)

	n := map[int]string{2: "xiaoli"}
	fmt.Println(n)
}
