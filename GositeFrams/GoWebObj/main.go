package main

import (
	"log"
	"net/http"
	"html/template"
)




var tpl *template.Template

func Home(w http.ResponseWriter, r *http.Request) {

	log.Print("Running Home Page....")

	WS := CreateTemplateStruct()


	tpl.ExecuteTemplate(w, "index.html", WS)

	return

}

func Page2(w http.ResponseWriter, r *http.Request) {

	log.Print("Running Page2....")
	
	WS := CreateTemplateStruct()

	// var listPowers []string

// 	User input 
	// SuperHero1 := SuperHuman{
	// 	Name: "Userinput",
	// 	Powers: listPowers,
	// }

// 	WS.B = append(SuperHero1)

	var BatPowers = []string{"Batman","Money","Intelligence"}

	WS.Body.Blist = append(BatPowers)
// 	WS.B = append(SuperHero1)


	tpl.ExecuteTemplate(w ,"Page2.html", WS)

}

func main() {

	tpl, _ = template.ParseGlob("./static/templates/*html")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))


	http.HandleFunc("/", Home)
	http.HandleFunc("/Page2", Page2)

	log.Println("Listening.....")
	log.Fatal(http.ListenAndServe(":8080", nil))


}