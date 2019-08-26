package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Println("Is string!" ,v)
	case int:
		fmt.Println("It is Map ", v)
	case []int:
		fmt.Println("It is Slice ", v)
	case map[string]int:
		fmt.Println("It'''''s is map 2v", v)
	default:
		fmt.Println("It is int ", v)
	}
}

func main() {
	a := make(map[string]int)
	do(a)
}
