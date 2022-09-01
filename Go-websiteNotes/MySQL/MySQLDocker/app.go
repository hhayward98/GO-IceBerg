package main

import (
    "database/sql"
    "fmt"
    "log"
     _ "github.com/go-sql-driver/mysql"
)

func main() {
	conn, err := sql.Open("mysql", "Test:toor@tcp(db:3306)/test_db")
	if err != nil {
		log.Fatal(err)
	}

	id  := 1 
	var name string

	if err := conn.QueryRow("SELECT name FROM test_tb WHERE id = ? LIMIT 1", id).Scan(&name); err != nil {
		log.Fatal(err)
	}


	fmt.Println(id, name)

}
