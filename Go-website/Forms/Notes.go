/*Forms*/

package main

import (
	"html/template"
	"net/http"
)

type LoginRequest struct {
	Username string
	Password string

}

func main() {
	tmpl := template.Must(template.ParseFiles("forms.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.methodPost {
			tmpl.Execute(w, nil)
			return
		}

		details := LoginRequest{
			Username: r.FormValue("Username"),
			Password: r.FormValue("Password"),
		}

		//First query the Username from database to see if it exist
		// if Username exists, Hash the input password and compair to the Password Hash that is currently stored with the Username in the database
		// if hash match then redirect to home page with user logged in.

		_ = details

		tmpl.Execute(W, struct{ Success bool }{true})
	})

	http.ListenAndServe(":8080", nil)
}