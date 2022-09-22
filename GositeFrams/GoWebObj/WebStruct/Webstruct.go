package WebStruct
// package main

import (
	"log"
	"net/http"
	"html/template"
)

// WebStruct is read into each template as the data.
// modify and add objects where needed within WebStruct

// Header, Body, footer are all arrays 

type WebStruct struct {
	Head Header
	Body body
	Foot Footer
}


// Domain name is used for routing links correctly on web app.
type Header struct {
	DomainName string	// http://example.com/8080 
	Style string	// /static/css/
	// Modify for what is needed.	
}


// an Array of somthing 
type body struct {
	Blist []string
}

// credits in the footer 
type Footer struct {
	Credits string
}



type SuperHuman struct {
	Name string
	Powers []string
}



func CreateTemplateStruct() WebStruct{
	H := Header{
		DomainName: "http://localhost:8080/",
		Style: "/static/css/",
	}

	var templist = []string{""}
	B := body{
		Blist: templist,
	}

	F := Footer{
		Credits: "Create By Aesir Constructs",
	}

	webstruct := WebStruct{
		Head: H,
		Body: B,
		Foot: F,
	}
	_ = webstruct

	return webstruct
}










// var tpl *template.Template

// func Home(w http.ResponseWriter, r *http.Request) {

// 	log.Print("Running Home Page....")

// 	WS := CreateTemplateStruct()


// 	tpl.ExecuteTemplate(w, "index.html", WS)

// 	return

// }

// func Page2(w http.ResponseWriter, r *http.Request) {

// 	log.Print("Running Page2....")
// 	// var listPowers []string
// 	WS := CreateTemplateStruct()


// // 	User input 
// 	// SuperHero1 := SuperHuman{
// 	// 	Name: "Userinput",
// 	// 	Powers: listPowers,
// 	// }



// 	var BatPowers = []string{"Batman","Money","Intelligence"}

// 	WS.Body.Blist = append(BatPowers)
// // 	WS.B = append(SuperHero1)


// 	tpl.ExecuteTemplate(w ,"Page2.html", WS)

// }

// func main() {

// 	tpl, _ = template.ParseGlob("./static/templates/*html")

// 	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))


// 	http.HandleFunc("/", Home)
// 	http.HandleFunc("/Page2", Page2)

// 	log.Println("Listening.....")
// 	log.Fatal(http.ListenAndServe(":8080", nil))


// }