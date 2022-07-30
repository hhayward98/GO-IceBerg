package main


import (
	"log"
	"net/http"
)

var userDB = map[string]string{
	"bob":	"password",
	"hunter": "password1",
}


var sessions = map[string]session{}

type session struct {
	username string
	expiry	time.Time
}

func (s session) isExpired() bool {
	return s.expiry.Before(time.Now())
}

func main() {

	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/secretPage" secretPage)

	log.Fatal(http.ListenAndServe(":8080", nil))
}