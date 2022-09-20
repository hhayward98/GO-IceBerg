package main

import (
	"log"
	"net/http"
	"html/template"

)

var tpl *template.Template


func Debugger(err error, Etype int) {
	if err != nil {
		if Etype == 1 {
			log.Fatal(err)
		}else if Etype == 2 {
			log.Print(err)
		}
	}
}




func Home(w http.ResponseWriter, r *http.Request) {
	log.Print("Running Home Page...")

	tpl.ExecuteTemplate(w, "index.html", nil)
	return
}




func AppRouts() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", Home)

	log.Fatal(http.ListenAndServe(":8088", nil))


}


func main() {



	tpl, err = template.ParseGlob("/static/templates/*html")
	Debugger(err, 2)

	log.Print("Listening.....")
	AppRouts()


}