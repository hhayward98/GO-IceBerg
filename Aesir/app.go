package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"

)

type InputForm struct {
	name string
	color string
	Favthing favthings
}

type favthings struct {
	T1 string
	T2 string
	T3 string
}

type InputForm struct{

	

}



func Debugger(err error) {
	if err != nil {
		log.Fatal(err)
	}
}


func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Running Home Page\n")

}

func Page2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Running Page2\n")

		//connect to database
	db, err := sql.Open("mysql", "Test:toor@(127.0.0.1:3308)/?parseTime=true")
	Debugger(err)

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("USE aesir")
	Debugger(err)

	log.Print("Connected to DB")
	



}

func AppRoutes() {

	//testing for docker container
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs)))
	
	http.HandleFunc("/", Home)
	http.HandleFunc("/Page2", Page2)
	log.Fatal(http.ListenAndServe(":8080", nil))

}


func main() {

	log.Print("Listening......")
	AppRoutes()

}