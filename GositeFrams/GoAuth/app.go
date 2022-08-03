package main 

import (
	"fmt"
	"net/http"
	"log"
	"github.com/google/uuid"
)

var tpl *template.Template



// use cach/database
var sessions = map[string]session{}


// session Values
type Session struct {
	Authenticated bool
	username string
	expiry time.Time 
}



func (s Session) Expired() bool {
	return s.expiry.Before(time.Now())
}



func login(w http.ResponseWriter, r *http.Request) {

	// run login stuff.....
	// sudo code
	// username from form == Uname
	Uname := "Bob"
	// 



	// if user is authorized

	// call following code once authorized
	seshToken := uuid.New()
	expiresAt := time.Now().Add(120 * time.Second)


	// add seshToken to session cach/database
	// since im not using cach/database
	// im doing somthing different
	sessions[seshToken] = Session{
		Authenticated: true,
		username: Uname,
		expiry: expiresAt,
	}


	//set cookie
	http.SetCookie(w, &http.Cookie{
		Name: "Session_token"
		Value: seshToken,
		Expires: expiresAt,
	})

	// redirect user to HomePage

}

func logout(w http.ResponseWriter, r *http.Request) {
	cook, err := r.Cookie("Session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			// redirect
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		// redirect
		return
	}

	seshToken := cook.Value

	delete(sessions, seshToken)

	http.SetCookie(w, &http.Cookie{
		Name: "Session_token",
		Value: "",
		Expires: time.Now(),
	})

	tpl.ExecuteTemplate(w, "login.html", "null")

}

func Home(w http.ResponseWriter, r *http.Request) {

	cook, err := r.Cookie("Session_token")
	if err != nil {
		if err == http.ErrNoCookie {

			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	seshToken := cook.Value

	Usesh, exists := Session[seshToken]
	if !exists {

		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if Usesh.Expired(){
		delete(sessions, seshToken)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	tpl.ExecuteTemplate(w, "Home.html", "auth")

}

func SeshRefresh(w http.ResponseWriter, r *http.Request) {

	cook, err := r.Cookie("Session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	seshToken := cook.Value

	Usesh, exists := sessions[seshToken]
	if !exists {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}


	NewSeshToken := uuid.NewString()
	expiresAt := time.Now().Add(120 * time.Second)


	sessions[NewSeshToken] = Session{
		Authenticated: true,
		username: Usesh.username,
		expiry: expiresAt,
	}

	delete(sessions, seshToken)

	http.SetCookie(w, &http.Cookie{
		Name: "Session_token",
		Value: seshToken,
		Expires: time.Now().Add(120 * time.Second),
	})


}


func main() {

	tpl, _ = template.ParseGlob("./static/Templates/*html")

	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/Home", Home)
	http.HandleFunc("/refresh", SeshRefresh)

	log.Print("Listening....")
	log.Fatal(http.ListenAndServe(":8080", nil))


}