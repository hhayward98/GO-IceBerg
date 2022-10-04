package main

import (

	"database/sql"
	"fmt"
	"github.com/lib/pq"

)


const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "<password>"
	dbname = "<dbname>"
)


func Check Error(err error) {
	if err != nil {
		panic(err)
	}
}



func main () {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// connect to database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close Database connection
	defer db.Close()

	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")

}



