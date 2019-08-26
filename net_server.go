package main

import (
	"net"
	"fmt"
)

func main(){
	ln ,err := net.Listen("tcp", ":8090")
	if err != nil {
		fmt.Println(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(conn)
		fmt.Println(&conn)
	}
	//go handleConnection(conn)
}
