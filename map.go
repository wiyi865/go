package main

import "fmt"

func main() {
	a := make(map[string]int)

	a["aa"] = 10
	a["bb"] = 190
	a["fs"] = 123
	//a =  {
	//	"aa": 10,
	//	"bb": 190,
	//	"fs":  123,
	//}

	for m, n := range a {
		fmt.Println(m, n)
	}
}
