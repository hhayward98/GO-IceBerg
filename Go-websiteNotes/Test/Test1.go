


package main

import (
	// "encoding/json"
	// "database/sql"
	"fmt"
	"log"
	// "time"
	"html/template"
	"net/http"
	"path/filepath"

	// "golang.org/x/crypto/bcrypt"
	// "github.com/gorilla/sessions"
	// "github.com/gorilla/mux"
)












// /*Sessions*/

// var (
// 	// replace with SHA256 Hash: 923ef6d931b1d39b9db72a224567f080f3360c493ac72295378fcb43e8cda6c2
// 	key = []byte("super-secret-key")
// 	store = sessions.NewCookieStore(key)
// )

// func secret(w http.ResponseWriter, r *http.Request) {
// 	session, _ := store.Get(r, "cookie-name")

// 	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
// 		http.Error(w, "Forbidden", http.StatusForbidden)
// 		// redirect to login page
// 		return
// 	}

// 	// else 
// 	// allow the user access. 
// }



// func logout(w http.ResponseWriter, r * http.Request) {
// 	session, _ := store.Get(r, "cookie-name")

// 	session.Values["authenticated"] = false
// 	session.Save(r, w)

// 	// redirect user to login page
// }




// func Index(w http.ResponseWriter, r *http.Request) {

// 	fs := http.FileServer(http.Dir("Static/"))
// 	http.Handle("/Home", http.StripPrefix("/Home", fs))


// }


// func Requesthandler() {

// 	r := mux.NewRouter()


// 	r.HandleFunc("/Home", Index)
// 	log.Fatal(http.ListenAndServe(":8080", r))


// }



type LoginRequest struct {
	Username string
	Password string
}


type RegisterDetails struct {
	Email string
	Username string
	Password string
	ConfPass string
}


func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 32)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func loginPage(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("static/Templates/Login.html"))

	data := LoginRequest{
		Username: r.FormValue("UserName"),
		Password: r.FormValue("PassWord"),

	}
	
	_ = data

	fmt.Println(data)

	//query username from database
	//hash password
	// check if password hash matches database 


	tmpl.Execute(w, struct{ Success bool }{true})
}

func Register(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/Templates/register.html"))

	data := RegisterDetails{
		Email: r.FormValue("email"),
		Username: r.FormValue("UserName"),
		Password: r.FormValue("PassWord"),
		ConfPass: r.FormValue("pass2"),
	}

	_ = data

	fmt.Println(data)

	tmpl.Execute(w, struct{ Success bool }{true})

}


func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/register", Register)

	log.Print("Listening....")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}