package main

import (
    "database/sql"
    "fmt"
    "log"
    // "time"
	"golang.org/x/crypto/bcrypt"

    _ "github.com/go-sql-driver/mysql"
)


func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 16)
	
	return string(bytes), err
}

func main() {
	db, err := sql.Open("mysql", "Test:toor@(127.0.0.1:3308)/?parseTime=true")

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("USE demon")
	if err != nil {
		log.Fatal(err)
	}


	var (
		Username string
		Password string
	)

	fmt.Println("Enter Username")
	fmt.Scanln(&Username)
	fmt.Println("Enter password")
	fmt.Scanln(&Password)

	// var (
	// 	UN string
	// 	PW string
		

	// )


	Hpass, _ := HashPassword(Password)

	fmt.Println(Hpass)

	// result, err := db.Exec(`INSERT INTO users (username, password, email, created_at) VALUES (?, ?, ?, ?)`, Username, Hpass, Email, time.Now())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// id, err := result.LastInsertId()
	// fmt.Println(id)

	// userCheck, _ := db.Query(`SELECT username, password, Email FROM users WHERE username = ?`, Username)

	// defer userCheck.Close()

	// for userCheck.Next() {
	// 	err := userCheck.Scan(&UN, &PW)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	}