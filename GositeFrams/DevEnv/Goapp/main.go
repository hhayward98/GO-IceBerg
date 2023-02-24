package main 

import (

	"log"
	"net/http"
	"html/template"
	"os"
	"github.com/joho/godotenv"

)



var tpl *template.Template


// Template Data Struct
type HTMLDATA struct {
	DomainName string
	Body string
	Foot string
}

// Init Data struct for templates
func HTMLDATASET() HTMLDATA{
	
	data := HTMLDATA{
		DomainName: "http://localhost:8080/",
		Body: "",
		Foot: "",
	}
	_ = data

	return data
}


// Debug Function for Error handling
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


// Home funtion
func Home(w http.ResponseWriter, r *http.Request) {

	HtmlData := HTMLDATASET()

	// Only allow GET method to access this page
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	tpl.ExecuteTemplate(w, "Home.html", HtmlData)

}


func AppRoutes() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))


	http.HandleFunc("/", Home)


	log.Fatal(http.ListenAndServe(":8080", nil))

}


func main() {

	tpl, _ = template.ParseGlob("./static/templates/*html")

	log.Println("Listening....")

	AppRoutes()

}
