package main 

import (
	"fmt"

)


func set(w http.ResponseWriter, r *http,Request) {
	http.SetCookie(w, &http.Cookie{
		Name: "Cookie",
		Value: "some value",
		Path: "/"
	})
}

func expire(w http.ResponseWriter, r *http,Request) {
	c, err := req.Cookie("session")
	
}

func login(w http.ResponseWriter, r *http,Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		id := make([]byte, 8)
		cookie = &http.Cookie{
			Name: "session",
			Value: id,String(),
			Secure: true,
			HttpOnly: true,
			Path: "/",
		}
		http.SetCookie(w, cookie)
	}
	fmt.Println(cookie)
}