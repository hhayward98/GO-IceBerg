package main 

import (
	// "fmt"
	"log"
	"net/http"
	"html/template"

)

var tpl *template.Template 

func Home(w http.ResponseWriter, r *http.Request) {
	log.Print("Running Home Page")

	tpl.ExecuteTemplate(w, "index.html", "")
	return
}

func main() {

	tpl, _ = template.ParseGlob("./static/Templates/*html")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", Home)

	log.Print("Listening.....")

	log.Fatal(http.ListenAndServe(":8080", nil))



}