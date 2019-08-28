package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	n := person{
		"王xiao刚",
		35,
		"上海市，China",
		"男",
	}

	fmt.Println(n)
	m, err := json.Marshal(n)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(m))
	p := &person{}
	erro := json.Unmarshal(m, p)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(*p)
}

type person struct {
	Name    string `json:"thename"`
	Age     int    `json:"age"`
	Address string `json:"address"`
	Sex     string `json:"性别"`
}
