package main


import {

	"database/sql"
	"fmt"
	"strings"
	"net/http"
	"html/template"
	"log"
	"time"
	"github.com/rivo/sessions"
	"github.com/rivo/users"
	"golang.org/x/crypto/bcrypt"
}


func MyHandler(w http.ResponseWriter, r *http.Request) {
	session, err := sessionStart(response, request, true)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(w, "we have session")
}

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("/static/index.html")
	tmpl.Execute(w, session)
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/sesh", MyHandler)

	log.Print("Listening......")
	log.Fatal(http.ListenAndServe(":8080", nil))
}


