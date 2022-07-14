package main

import (
	"log"
	"fmt"
	"html/template"
	"net/http"


)

type LoginRequest struct {
	Username string
	Password string
}


type RegisterDetails struct {
	Email string
	Username string
	Password string
	ConfPass string
}

func loginPage(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("static/Templates/Login.html"))

	data := LoginRequest{
		Username: r.FormValue("UserName"),
		Password: r.FormValue("PassWord"),

	}
	
	_ = data

	fmt.Println(data)

	tmpl.Execute(w, struct{ Success bool }{true})
}

func Register(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/Templates/register.html"))

	data := RegisterDetails{
		Email: r.FormValue("email"),
		Username: r.FormValue("UserName"),
		Password: r.FormValue("PassWord"),
		ConfPass: r.FormValue("pass2"),
	}

	_ = data

	fmt.Println(data)

	tmpl.Execute(w, struct{ Success bool }{true})

}


func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// need to change so when you click login button it navigates http://localhost:8080/login rather than the actaul source file

	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/register", Register)

	log.Print("Listening....")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}