package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	b, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Print(err)
	}
	//fmt.Println(b)
	str := string(b)
	fmt.Println(str)
}
