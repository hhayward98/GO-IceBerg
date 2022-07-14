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
	"github.com/gorilla/sessions"
	"github.com/gorilla/mux"
)


/*Forms*/


type LoginRequest struct {
	Username string
	Password string
}

type RegisterDetails struct {
	Email string
	Username string
	Password string
}


/*Sessions*/

var (
	// replace with SHA256 Hash: 923ef6d931b1d39b9db72a224567f080f3360c493ac72295378fcb43e8cda6c2
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
	session.Values["username"] = // set to input username 
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


/*Logging*/
// not really implemented for now
func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}


func TestLogin(w http.ResponseWriter, r *http.Request) {
	// connect to DB 
	db, err := sqlOpen("mysql", "root:root@(127.0.0.1:3306)/root?parseTime=true")
	if err != ni;; {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}


	tmpl := template.Must(template.ParseFiles("Static/Template/Login.html"))
	session, _ := store.Get(r, "cookie-name")

	if r.Method != http.methodPost {
		tmpl.Execute(w, nil)
		return
	}

	details := LoginRequest{
		// get input from form
		username: r.FormValue(UserName)
		password: r.FormValue(PassWord)
	}



	query, err := db.Query(`SELECT username FROM users WHERE username = ?`, details.username)
	if err != nil {
		log.Fatal(err)
	}

	if details.username != query {
		// user dose not exist 
		//route to register page
	} else {
		query, err := db.Query(`SELECT password FROM users WHERE username = ?`, details.username)
		if err != nil {
			log.Fatal(err)
		}
		check := CheckPasswordHash(details.password, query)
		if check = true {
			session.Values["authenticated"] = true
			session.Values["username"] = username
			session.Save(r, w)
		} else {
			session.Values["authenticated"] = false
			// flash invalid password
			// refresh page or wrap code in loop 
			// 
		}
	}

	// _ = details

	tmpl.Execute(w, struct{ Success bool }{true})
}



/*Main*/

func main() {

	// init router
	r := mux.NewRouter()

	// Route Handlers 
	// r.HandleFunc("Static/index.html", HomePage).Methods("GET")
	// r.HandleFunc("Static/Template/Login.html", Login).Methods("POST")
	// r.HandleFunc("Static/Template/register.html", Register).Methods("POST")
	// r.HandleFunc("Static/Template/Login.html", Logout).Methods("POST")

	http.ListenAndServe(":8080", nil)
}