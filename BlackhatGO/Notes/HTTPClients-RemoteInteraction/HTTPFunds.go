package main

import (
	"encoding/json"
	"log"
	"net/http"
)


func CallingAPI() {

	req1, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatalln(err)
	}

	defer req1.Body.Close()


	form := url.Values{}
	form.Add("Whats", "Good")

	// using PostForm 
	req3, err = http.PostForm("https://www.google.com/robots.txt", form)
	if err != nil {
		log.Fatalln(err)
	}
}


// contains expected elements from server response
type Status struct {
	Message string
	Status string
}


func main() {
	// sends Post request and decodes response
	res, err := http.Post("http://IP:PORT/ping","application/json", nil)
	if err != nil {
		log.Fatalln(err)
	}

	var status Status
	// Decoding response Body
	if err := json.NewDecoder(res.Body).Decode(&status); err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()
	log.Printf("%s ->\n", status.Status, status.Message)

}


