package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for _, i := range os.Args[1:] {
		s += sep + i
		sep = " "
	}
	fmt.Println(s)
}
