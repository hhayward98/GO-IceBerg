/*Initalizing Database to use with web app*/

package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//connect to database
	db, err := sql.Open("mysql", "Test:toor@(127.0.0.1:3306)/DDsql?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// create tables

	{
		query:= `
			CREATE TABLE users (
				id INT AUTO_INCREMENT,
				username TEXT NOT NULL,
				password TEXT NOT NULL,
				email TEXT NOT NULL,
				created_at DATETIME,
				PRIMARY KEY (id)
			);`

		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}

	// create Admin and insert into database
	{
		username := "Admin"
		password := "Admin"
		email := "Admin@boss.com"
		createdAt := time.Now()

		result, err := db.Exec(`INSERT INTO users (username, password, email, created_at) Values(?, ?, ?)`, username, password, email, createdAt)
		if err != nil {
			log.Fatal(err)
		}

		id, err := result.LastInsertId()
		fmt.Println(id)
	}
}