package main 

import (

	"net/http"
	"log"
	"html/template"
)

var tpl *template.Template

func Method1(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("static/Templates/Method1.html"))

	tmpl.Execute(w, nil)
}

func Method2(w http.ResponseWriter, r *http.Request) {

	tpl.ExecuteTemplate(w, "Method2.html", "auth")
}



func main() {

	tpl, _ = template.ParseGlob("./static/Templates/*html")

	http.HandleFunc("/Method1", Method1)
	http.HandleFunc("/Method2", Method2)

	log.Print("Listening....")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
