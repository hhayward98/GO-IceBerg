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


func Imgrotate(w http.ResponseWriter, r *http.Request) {
	log.Print("Running Image Rotate.....")

	tpl.ExecuteTemplate(w, "ImgRotate.html", "")

}

func TempConv(w http.ResponseWriter, r *http.Request) {
	log.Print("Running TempConv.....")

	tpl.ExecuteTemplate(w, "TempConv.html", "")


}

func main() {

	tpl, _ = template.ParseGlob("./static/templates/*html")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", Home)
	http.HandleFunc("/ImageRotate", Imgrotate)
	http.HandleFunc("/TempConvert", TempConv)

	log.Print("Listening.....")

	log.Fatal(http.ListenAndServe(":8080", nil))



}