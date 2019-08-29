package main

/*"""
从一个文件中拷贝内容到另外一个文件
"""
*/

import (
	//"fmt"
	"io"
	"log"
	"os"
)

func copyfile(fromfile, tofile string) error {
	sfile, err := os.Open(fromfile)
	if err != nil {
		log.Fatal("In Function copyfile", err)
	}

	tfile, err := os.Create(tofile)
	//tfile, err := os.OpenFile(tofile, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("In Function copyfile", err)
	}
	_, error := io.Copy(tfile, sfile)
	return error
}

func main() {
	filea := os.Args[1]
	fileb := os.Args[2]
	err := copyfile(filea, fileb)
	if err != nil {
		log.Fatal(err)
	}
}
