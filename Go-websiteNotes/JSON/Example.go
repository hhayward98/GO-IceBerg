/*Json*/

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Fisrtname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Age 	int 	`json:"age"`
}

func main() {
	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
		bar user User
		json.NewDecoder(r.Body).Decode(&user)

		fmt.Fprintf(w, "%s %s is %d years old!", user.Fisrtname, user.Lastname, user.Age)

	})

	http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
		peter := User{
			Fisrtname: "junter",
			Lastname: "Hayward",
			Age: 	  69
		}

		json.NewEncoder(w).Encode(peter)
	})

	http.ListenAndServe(":8080", nil)
}