package main

import (
	"fmt"
)

func count(out chan<- int) {
	for i := 0; ; i++ {
		out <- i
	}
	close(out)
}

func square(in <-chan int, out chan<- int) {
	for num := range in {
		out <- num * num
	}
	close(out)
}

func do_print(in <-chan int) {
	for result := range in {
	fmt.Println(result)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	go count(naturals)
	go square(naturals, squares)
	do_print(squares)
}
