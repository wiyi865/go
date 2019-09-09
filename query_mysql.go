package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"log"
)

func main(){
	age := 27
	db, err := sql.Open("mysql", "go:go@/go")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SElECT name FROM users WHERE age = ?", age)

	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s is %d\n", name, age)
	}
	if err := rows.Err(); err != nil{
		log.Fatal(err)
	}
}
