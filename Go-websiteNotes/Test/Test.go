/*Forms*/

package main

import (
	"fmt"
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
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		data := LoginRequest{
			Username: r.FormValue("Username"),
			Password: r.FormValue("Password"),
		}

		_ = data

		if data.Username != "" {
			if data.Password != ""{
				tmpl.Execute(w, struct{ Success bool }{true})
				fmt.Println(data.Username)
				fmt.Println(data.Password)
			}else {
				fmt.Println("Password cant be null")
				tmpl.Execute(w, struct{ Success bool }{false})
			}
		}else {
			fmt.Println("Username cant be null")
			tmpl.Execute(w, struct{ Success bool }{false})
		}

	})

	http.ListenAndServe(":8080", nil)
	// navigate to http://localhost:8080/
}