package main 

import (

	"net/http"
	"log"
	"html/template"
)

var tpl *template.Template

func Method1(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("static/templates/Method1.html"))

	tmpl.Execute(w, nil)
}

func Method2(w http.ResponseWriter, r *http.Request) {

	tpl.ExecuteTemplate(w, "Method2.html", "auth")
}

func Method3(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "static/templates/Method3.html", 201)
	return
}

func main() {

	tpl, _ = template.ParseGlob("./static/Templates/*html")

	http.HandleFunc("/Method1", Method1)
	http.HandleFunc("/Method2", Method2)
	http.HandleFunc("/Method3", Method3)

	log.Print("Listening....")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
