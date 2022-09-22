package Test52

import (
	"log"
	"time"
	"net/http"

)

var sessions = map[string]session{}

type session struct {
	username string
	expiry	time.Time
}


func (s Session) Expired() bool {
	return s.expiry.Before(time.Now())
}




func login(w http.ResponseWriter, r *http.Request) {




	// generate new session token id
	seshToken := "randString" // generate randome string
	expiresAt := time.Now().Add(120 * time.Second)

	// set tokenid in session map with its information
	sessions[seshToken] = session{
		username: userinfo.username,
		expiry: expiresAt,
	}


	// set client cookie 
	http.SetCookie(w, &http.Cookie{
		Name: "Session_token",
		Value: seshToken,
		Expires: expiresAt,
	})

}





func Refresh(w http.ResponseWriter, r *http.Request) {
	cook, err := r.Cookie("session_token")
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
	if Usesh.Expired() {
		delete(sessions, seshToken)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// if previous session is valid make a new session token for current user
	NewSeshToken := "randString1"
	expiresAt := time.Now().Add(120 * time.Second)

	sessions[NewSeshToken] = session{
		username: Usesh.username,
		expiry: expiresAt,
	}

	// delete old session token
	delete(sessions, seshToken)

	// set new token
	http.SetCookie(q, &http.Cookie{
		Name: "session_token",
		Value: NewSeshToken,
		Expires: time.Now().Add(120 * time.Second),
	})

}



func TestPage(w http.ResponseWriter, r *http.Request) {
	log.Print("Running TestPage....")




}