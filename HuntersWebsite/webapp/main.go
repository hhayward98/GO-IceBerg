package main

import (
	"log"
	"net/http"
	"html/template"
	// "github.com/HuntersRoutes/WebsiteRoutes"
)

var tpl *template.Template

type HTMLDATA struct {
	DomainName string
	Body string
	Foot string
}

func Debugger(err error, Etype int) {

	if err != nil {
		// Error type 1 will end the program
		if Etype == 1 {
			log.Fatal(err)
		}else if Etype == 2 { // Error type 2 returns error without killing app
			log.Print("===========================")
			log.Print(err)
			log.Print("===========================")
		}
	}
}


func HTMLDATASET() HTMLDATA{
	data := HTMLDATA{
		DomainName: "http://localhost",
		Body: "",
		Foot: "",
	}
	_ = data

	return data
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("Running Health Check......")

	// Run Test: 
	// Check response for each route 
	// ....

	// if no errors are thrown then return response


	return

}


func Home(w http.ResponseWriter, r *http.Request) {

	HtmlData := HTMLDATASET()

	// Only allow GET method to access this page
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	tpl.ExecuteTemplate(w,"index.html", HtmlData)

}


func Skills(w http.ResponseWriter, r *http.Request) {

	HtmlData := HTMLDATASET()

	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	tpl.ExecuteTemplate(w, "skills.html", HtmlData)
}


func projects(w http.ResponseWriter, r *http.Request) {

	HtmlData := HTMLDATASET()

	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	tpl.ExecuteTemplate(w, "projects.html", HtmlData)


}



func AppRoutes() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", Home)
	http.HandleFunc("/Skills", Skills)
	http.HandleFunc("/Projects", projects)


	log.Fatal(http.ListenAndServe(":8080", nil))

}



func main() {

	tpl, _ = template.ParseGlob("static/templates/*html")


	log.Println("Listening......")

	AppRoutes()
	
}