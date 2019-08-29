package main

import (
	"fmt"
	//"os"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	go func() {
		for i := 0; ; i++ {
			//for i := 0; i < 10; i++ {
			naturals <- i
		}
		//	os.Exit(1)
	}()
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()
	for {
		fmt.Println(<-squares)
	}
}
