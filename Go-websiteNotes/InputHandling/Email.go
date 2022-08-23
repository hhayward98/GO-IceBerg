package main

import (
	"fmt"
	"log"
	"net/http"
	"net/mail"
)

func validateEmail(addy string) (string, bool) {
	addr, err := mail.ParseAddress(addy)
	if err != nil {
		return "", false

	}
	return addr.Address, true
}

func Home(w http.ResponseWriter, r *http.Request) {
	log.Print("Running Home Page")


	
}

func main() {

	http.HandleFunc("/", Home)
	log.Print("Listening.....")
	log.Fatal(http.ListenAndServe(":8080", nil))
}