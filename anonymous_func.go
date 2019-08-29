package main

import "fmt"

func squares() func() int {
	var x int
	fmt.Println("x的值在最外层为: ", x)
	return func() int {
		fmt.Println("x的值在最内层为: ", x)
		x++
		return x * x
	}
}
func main() {
	f := squares()
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
	fmt.Println(f()) // "16"
}
