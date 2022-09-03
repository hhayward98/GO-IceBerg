package main 

import (
	"fmt"
	"log"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Debugger(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func DBConnect(w http.ResponseWriter, r *http.Request) {
	log.Print("Connecting to Database")


	db, err := sql.Open("mysql", "test:toor@tcp(db:3306)/sqldock")
	Debugger(err)
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("USE sqldock")
	Debugger(err)

	fmt.Println("Connection Success!!!")

}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home Page")

}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/DBconnect", DBConnect)
	log.Print("Listening....")
	log.Fatal(http.ListenAndServe(":8088", nil))

}