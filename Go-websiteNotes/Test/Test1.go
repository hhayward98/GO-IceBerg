package main


import (
	// "database/sql"
	// "fmt"
	// "log"
	// "time"
	"html/template"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	_ "github.com/go-sql-driver/mysql"
)

var (
	// replace with SHA256 Hash: 923ef6d931b1d39b9db72a224567f080f3360c493ac72295378fcb43e8cda6c2
	key = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func TestLogin(w http.ResponseWriter, r *http.Request) {
	// connect to DB 
	// db, err := sqlOpen("mysql", "Test:Test@(127.0.0.1:3306)/Test?parseTime=true")
	// if err != ni;; {
	// 	log.Fatal(err)
	// }
	// if err := db.Ping(); err != nil {
	// 	log.Fatal(err)
	// }


	tmpl := template.Must(template.ParseFiles("Static/Template/Login.html"))
	session, _ := store.Get(r, "cookie-name")

	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	details := UserLogin{
		// get input from form
		Username: r.FormValue("Username"),
		Password: r.FormValue("Password"),
	}



	// query, err := db.Query(`SELECT username FROM users WHERE username = ?`, details.username)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if details.username != query {
	// 	// user dose not exist 
	// 	//route to register page
	// } else {
	// 	query, err := db.Query(`SELECT password FROM users WHERE username = ?`, details.username)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	check := CheckPasswordHash(details.password, query)
	// 	if check = true {
	// 		session.Values["authenticated"] = true
	// 		session.Values["username"] = username
	// 		session.Save(r, w)
	// 	} else {
	// 		session.Values["authenticated"] = false
	// 		// flash invalid password
	// 		// refresh page or wrap code in loop 
	// 		// 
	// 	}
	// }

	

	tmpl.Execute(w, struct{ Success bool }{true})
}




type UserLogin struct {
	Username string
	Password string
}



func main() {

	r := mux.NewRouter()

	// Route Handlers 
	r.HandleFunc("login.html", TestLogin).Methods("POST")

	http.ListenAndServe(":8080", nil)

}