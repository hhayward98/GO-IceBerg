package main

import (
	// "fmt"
	"log"
	"net/http"
	"html/template"
)

var tpl *template.Template

type Sesh struct {

	Authenticated bool
	User string
	Power SuperPower
}

type SuperPower struct{
	PowerID int64
	Name string
	ability string

}

func Home(w http.ResponseWriter, r *http.Request) {

	obj1 := SuperPower{01,"Flying", "mobility"}
	obj2 := Sesh{true, "hunter", obj1}

	tpl.ExecuteTemplate(w, "first.html", obj2)
}

func Page2(w http.ResponseWriter, r *http.Request) {

	tpl.ExecuteTemplate(w, "second.html", "null")
}

func main() {

	tpl, _ = template.ParseGlob("./templates/*html")

	http.HandleFunc("/", Home)
	http.HandleFunc("/page2", Page2)

	log.Print("Listening......")
	log.Fatal(http.ListenAndServe(":8080", nil))

}