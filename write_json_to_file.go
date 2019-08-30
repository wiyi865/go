package main

import (
	"fmt"
	"log"
	"encoding/json"
	"os"
)

type person struct {
	Name string `json:"name"`
	Age int `"json:age"`
	Sex string `"json:sex"`
}



func main(){
	p := person {
	"xiaoli",
	20,
	"女",
	}
	content, err := json.Marshal(&p)
	if err != nil {
		log.Fatal("json.Marshal Err : ", err)
	}

	fd, err := os.Create("/tmp/1.txt")
	if err != nil {
		log.Fatal("os.Create Err : ", err)
	}
	length, err := fd.Write(content)

	if err != nil {
		log.Fatal("write  Err : ", err)
	}

	fmt.Println( length , "字节 has been wroten!")
}
