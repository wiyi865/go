package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"log"
)

func main(){
	//db, err := sql.Open("mysql", "root:root@/go")
	//db, err := sql.Open("mysql", "root:@/go")
	db, err := sql.Open("mysql", "go:go@/go")
	if err != nil {
		log.Fatal(err)
	}
	result, err := db.Exec(
		"INSERT INTO users (name, age) VALUES (? ,?)",
		"gopher",
		"27",
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.RowsAffected())
}
