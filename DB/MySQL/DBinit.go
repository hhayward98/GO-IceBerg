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

type SuperHuman struct {
	Name string
	PassiveP string
	AttackP string
}


func main() {

	//connect to database
	db, err := sql.Open("mysql", "Test:toor@(127.0.0.1:3308)/?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("USE superhumans")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to Database")



	supervillains := make([]*SuperHuman,0)

	VillainRows, _ := db.Query(`SELECT villainname, passivepower, attackpower FROM villains`)
	if err != nil {
		log.Fatal(err)
	}

	for VillainRows.Next() {
		Villain := new(SuperHuman)
		
		if err := VillainRows.Scan(&Villain.Name, &Villain.PassiveP, &Villain.AttackP); err != nil {
			panic(err)
		}
		
		supervillains = append(supervillains, Villain)

	}

	if err := VillainRows.Err(); err != nil {
		panic(err)
	}

	VillainRows.Close()




	log.Println(supervillains)
	for i, s := range supervillains {
		fmt.Println(i,s)
		Name := s.Name
		fmt.Println(Name)
	}



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