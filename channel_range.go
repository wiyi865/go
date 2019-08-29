package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	go func() {
		for i := 0; ; i++ {
			naturals <- i
		}
	}()
	go func() {
		for x := range naturals {
			squares <- x * x
		}
	}()
	for y := range squares {
		fmt.Println(y)
	}
}
