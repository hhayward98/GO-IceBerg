package main


import (
	"log"
	"net/http"
)

var userDB = map[string]string{
	"bob":	"password",
	"hunter": "password1",
}

// stores the users session 
//using database or cache to store sessions for Large scale applications
var sessions = map[string]session{}


// session variables 
type session struct {
	username string
	expiry	time.Time
}



func (s session) Expired() bool {
	return s.expiry.Before(time.Now())
}

// struct used to model request body
type Testmodel struct {
	username string `json:"username"`
	password string `json:"password"`
}


func login(w http.ResponseWriter, r *http.Request) {
	var userinfo Testmodel

	err := json.NewDecoder(r.Body).Decode(&userinfo)
	if err != nil {
		// return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// geting password 
	pass, ok := users[userinfo.username]


	// if password exists for the given user and is the same as stored then continue 
	// else return unauthorized 
	if !ok || pass != userinfo.password {
		w.WriteHeader(http.StatusUnauthorized)
		return 
	}

	// generate new session token id
	seshToken := "randString"
	expiresAt := time.Now().Add(120 * time.Second)

	// set tokenid in session map with its information
	sessions[seshToken] = session{
		username: userinfo.username,
		expiry: expiresAt,
	}


	// set client cookie 
	http.SetCookie(w, &http.Cookie{
		Name: "Session_token"
		Value: seshToken,
		Expires: expiresAt,
	})

}

func logout(w http.ResponseWriter, r *http.Request) {
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


	// delete user session from session map
	delete(sessions, seshToken)

	// letting the client know the cookie is expired by seting Value to empty and expiry to time.Now()
	http.SetCookie(w, &http.Cookie{
		Name: "session_token",
		Value: "",
		Expires: time.Now(),
	})
}


func secretPage(w http.ResponseWriter, r *http.Request) {
	// get session token from request cookies
	cook, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			//when no cookie is set return unauthorized
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// return bad request for gen error 
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	seshToken := cook.Value

	// get Session from session map
	Usesh, exists := session[seshToken]
	if !exists {
		// if sessiontoken dose not exists, return unauthorized
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// when session expired we can delete session and return unauthorized 
	if Usesh.Expired() {
		delete(sessions, seshToken)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}


	// If session is valid 
	// execute page 
}


// Refresh Session Tokens for user to keep them logged in

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

func main() {

	http.HandleFunc("/login", login)
	http.HandleFunc("/secretPage" secretPage)
	http.HandleFunc("/refresh", Refresh)
	http.HandleFunc("/logout", logout)

	log.Fatal(http.ListenAndServe(":8080", nil))
}


