package main

import (
	"fmt"
	"log"
)


func Home(w http.ResponseWriter, r *http.Request) {
	log.Print("Running Home Page\n")

}


func main() {

	http.HandleFunc("/", Home)
	log.Print("Listening....")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

