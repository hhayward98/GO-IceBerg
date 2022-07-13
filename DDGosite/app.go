package main

import (
	"encoding/json"
	"database/sql"
	"fmt"
	"log"
	"time"
	"html/template"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"github.com/gorrilla/sessions"
)



/*Sessions*/

var (
	// replace with a 32 byte AES-256 key
	key = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func secret(w http.ResponseWriter, r *http.Request) {
	session, _ : = store.Get(r, "cookie-name")

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		// redirect to login page
		return
	}

	// else 
	// allow the user access. 
}

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")


	// Authenticate user
	// recieve Username and password from login form as input
	// connect to database and check to see if username exist 
	// if username exists then run HashPassword() func on the input password
	// if the Hash of the input password Matches the Hashed password that is stored in the database 
	// Then we Authenticate the user


	// set user as Authenticated
	session.Values["authenticated"] = true
	session.Save(r, w)

	// Then redirect the user to the Home page with them logged in


}

func logout(w http.ResponseWriter, r * http.Request) {
	session, _ := store.Get(r, "cookie-name")

	session.Values["authenticated"] = false
	session.Save(r, w)

	// redirect user to login page
}





/*Password Hashing*/

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 32)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}


/*Loggin*/
// not really implemented for now
func loggin(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}




