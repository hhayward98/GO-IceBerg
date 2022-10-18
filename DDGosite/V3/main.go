package main

import (

	"log"
	"strings"
	"time"
	"net/http"
	"net/mail"
	"html/template"

	"database/sql"

	"golang.org/x/crypto/bcrypt"

	_ "github.com/go-sql-driver/mysql"


)

var tpl *template.Template


type LoginDetails struct {
	Username string
	Password string
}


type RegisterDetails struct {
	Email string
	Username string
	Password string
	ConfPass string
}


type HTMLDATA struct {
	Header string
	Body string
	Foot string
}





func AppRoutes() {


	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	

	http.HandleFunc("/", Home)



	http.HandleFunc("/Login", Login)
	http.HandleFunc("/Register", Register)
	http.HandleFunc("/Logout", logout)
	http.HandleFunc("/SecretPage", secretPage)

	http.HandleFunc("/Tools", ToolsPage)
	http.HandleFunc("/Projects", Projects)
	http.HandleFunc("/DemonDAO", DAOPage)
	http.HandleFunc("/Services", Services)


	log.Fatal(http.ListenAndServe(":8088", nil))

	//TLS
	// err := http.ListenAndServeTLS(":9000", "localhost.crt", "localhost.key", nil)
	// Debugger(err, 1)

}


func main() {

	tpl, _ = template.ParseGlob("./static/templates/*html")

	log.Print("Listening....")

	AppRoutes()

}