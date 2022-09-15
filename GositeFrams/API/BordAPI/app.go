package main

import (
	"log"
	"net/http"
	"html/template"


)

var tpl *template.Template

func Home(w http.ResponseWriter, r *http.Request) {
	log.Print("Running Home Page")

	// BoardAPI.Add(5,4)

	tpl.ExecuteTemplate(w, "index.html", "")

}


func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", Home)

	log.Print("Listening.....")

	log.Fatal(http.ListenAndServe(":8080", nil))

}