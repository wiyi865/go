package main

import (
	"fmt"
	//"os"
	"sync"
)

var (
	g        sync.WaitGroup
	naturals = make(chan int)
	squares  = make(chan int)
)

func num() {
	defer g.Done()
	for i := 0; i < 100; i++ {
		naturals <- i
	}
}

func doSquare() {
	defer g.Done()
	for x := range naturals {
		squares <- x * x
	}
}
func main() {

	g.Add(1)
	go num()

	g.Add(1)
	go doSquare()

	for y := range squares {
		fmt.Println(y)
	}
	g.Wait()
}
