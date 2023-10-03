package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func init() {

	_, err := sql.Open("sqlite3", "models/main.db")

	if err != nil {
		fmt.Println(err.Error())
	}

}

func main() {

	fmt.Println("Hi")

}
