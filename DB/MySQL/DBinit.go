package main 

import (
	"database/sql"
	"fmt"
	"log"
	// "time"
	// "golang.org/x/crypto/bcrypt"
	_ "github.com/go-sql-driver/mysql"
)

// func HashPassword(password string) (string, error) {
// 	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 16)
// 	return string(bytes), err
// }

func main() {

	//connect to database
	db, err := sql.Open("mysql", "Test:toor@(127.0.0.1:3308)/?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("USE pytest")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to Database")


	// user table 
	// {
	// 	query:= `
	// 		CREATE TABLE users (
	// 			id INT AUTO_INCREMENT,
	// 			username TEXT NOT NULL,
	// 			password TEXT NOT NULL,
	// 			email TEXT NOT NULL,
	// 			created_at DATETIME,
	// 			PRIMARY KEY (id),
	// 		);`

	// 	if _, err := db.Exec(query); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	// {
	// 	username := "Admin"
	// 	password, _ := HashPassword("Admin")
	// 	email := "Admin@boss.com"
	// 	createdAt := time.Now()

	// 	result, err := db.Exec(`INSERT INTO users (username, password, email, created_at) Values(?, ?, ?, ?)`, username, password, email, createdAt)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	id, err := result.LastInsertId()
	// 	fmt.Println(id)
	// }
}