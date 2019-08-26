package main

import (
	"fmt"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"net/http"
	"os"
)

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

func do_get(url string) {
	s, err := http.Get(url)
	//s, err := http.Get("www.baidu.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	url_content, err := ioutil.ReadAll(s.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	s.Body.Close()
	//fmt.Println(ss)
	//fmt.Println(string(ss))
	result := ConvertToString(string(url_content), "gbk", "utf-8")
	fmt.Println(string(result))
	//fmt.Printf("%s",ss)
}

func main() {
	lists := os.Args[1:]
	if len(lists) == 0 {
		url := "https://www.163.com"
		do_get(url)
	} else {
		for _, url := range os.Args[1:] {
			do_get(url)
		}
	}
}
