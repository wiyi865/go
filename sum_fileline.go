/*
读取/tmp/1.txt文件内容，并求和
1
2
7
90
800
677
需要注意的是读取每行结果为[]byte,转化为string再转化为Int,暂时不确定是不是正确的办法，当前已实现累计求和。
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	var sum int
	f := "/tmp/1.txt"
	fd, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	read := bufio.NewReader(fd)
	for {
		line, _, err := read.ReadLine()
		if err == io.EOF {
			break
		}
		//fmt.Println(string(line))
		int_line, err := strconv.Atoi(string((line)))
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(int_line)
		sum += int_line
	}
	fmt.Println(sum)
}
