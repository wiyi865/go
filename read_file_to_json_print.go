/*
/tmp/1.txt内容为:
{
	"name":"xiaoli",
	"Age":20,
	"Sex":"女"
}
*/
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
	fd, err := os.Open("/tmp/1.txt")
	if err != nil {
		log.Fatal("os.Open Err : ", err)
	}
	var p *person
	o := json.NewDecoder(fd).Decode(&p)
	if o != nil {
		log.Fatal("json.newdecoder : ", o)
	}

		fmt.Println(p.Age)
	//for _, one := range p {
	//	fmt.Println(one.Age)
	//}
}
