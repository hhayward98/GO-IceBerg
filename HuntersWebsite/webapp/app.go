package main

import (
	"fmt"
	"net/http"
	"log"
)

var tpl *template.Template

var Ptpl *template.Template

func Home(w http.ResponseWriter, r *http.Request) {

	tpl.ExecuteTemplate(w,"index.html", "null")

}


func Skills(w http.ResponseWriter, r *http.Request) {

	tpl.ExecuteTemplate(w, "skills.html")
}


func projects(w http.ResponseWriter, r *http.Request) {

	tpl.ExecuteTemplate(w, "projects.html", "null")
}

func EGEN(w http.ResponseWriter, r *http.Request) {

	tpl.ExecuteTemplate(w, "EGEN310.html", "null")
}


func EmbeddedSystems(w http.ResponseWriter, r *http.Request) {

	tpl.ExecuteTemplate(w, "EmbeddedSystems.html", "null")
}


func DemonDAO(w http.ResponseWriter, r *http.Request) {

	tpl.ExecuteTemplate(w, "DemonDAO.html", "null")
}

func RDP(w http.ResponseWriter, r *http.Request) {

	Ptpl.ExecuteTemplate(w, "RecDecPar.html", "null")
}


func main() {

	tpl, _ = template.ParseGlob("static/templates/*html")

	Ptpl, _ = template.ParseGlob("static/templates/Proj/*html")


	http.HandleFunc("/", Home)
	http.HandleFunc("/Skills", Skills)
	http.HandleFunc("/Projects", projects)
	http.HandleFunc("/EmbeddedSystems", EmbeddedSystems)
	http.HandleFunc("/EGEN310", EGEN)
	http.HandleFunc("/DemonDAO", DemonDAO)
	http.HandleFunc("/RDP", RDP)

	log.Print("Listening......")

	log.Fatal(http.ListenAndServe(":8080", nil))
	
}