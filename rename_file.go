package main

import (
	"fmt"
	"os"
)

func main() {
	file := os.Args[1]
	err := os.Rename(file, "1.rename.txt")
	if err != nil {
		fmt.Println("not ok")
	} else {
		fmt.Println("ok")
	}
}
