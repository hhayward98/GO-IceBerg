package main

import (
	"log"
	"fmt"
	"html/template"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"github.com/gorilla/sessions"


)

var (

	key = []byte("923ef6d931b1d39b9db72a224567f080f3360c493ac72295378fcb43e8cda6c2")
	store = sessions.NewCookieStore(key)
	// user = nil
	// store = sessions.NewCookieStore(user)

)

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

	session, _ := store.Get(r, "cookie-name")

	data := LoginRequest{
		Username: r.FormValue("UserName"),
		Password: r.FormValue("PassWord"),

	}
	
	_ = data

	fmt.Println(data)

	//query username from database
	// temp name untill I link DB
	name := "Temp"
	Pass := "$2a$14$vk9UXe4zniQ/kZJM/INvr.P1LbE6EMjuISqVGxFOk0m4mqU0pp.OS"


	// Hpass, _ := HashPassword(data.Password)

	// check if password hash matches database 

	Match := CheckPasswordHash(data.Password, Pass)

	if Match == true {

		session.Values["authenticated"] = true
		session.Values["User"] = name
		session.Save(r, w)

		// redirect user to secretpage

		tmpl := template.Must(template.ParseFiles("static/Templates/secretPage.html"))
		tmpl.Execute(w, nil)
		return

	} else if Match == false {
		// flash error message
		fmt.Println("Invalid Login")
	}


	tmpl.Execute(w, struct{ Success bool }{true})
}

func Register(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/Templates/register.html"))

	session, _ := store.Get(r, "cookie-name")

	data := RegisterDetails{
		Email: r.FormValue("email"),
		Username: r.FormValue("UserName"),
		Password: r.FormValue("PassWord"),
		ConfPass: r.FormValue("pass2"),
	}

	_ = data

	fmt.Println(data)

	HashPass, _ := HashPassword(data.Password)
	// check password matches or build front end to handle that.
	Match := CheckPasswordHash(data.ConfPass, HashPass)


	Userstatus := true
	/*
		I will need to change the authentication to account for existing usernames or emails
		for now it never fails to register someone.

	*/



	if Userstatus == true {
		// both passwords they input match 

		session.Values["authenticated"] = true
		session.Values["User"] = data.Username
		session.Save(r, w)

		// redirects user to secret page when they register
		tmpl := template.Must(template.ParseFiles("static/Templates/secretPage.html"))
		tmpl.Execute(w, nil)
		return

	} else if Userstatus == false {
		// printing to terminal not webpage
		fmt.Println("Passwords did not match")

	}





	// Insert New user into DataBase 




	tmpl.Execute(w, struct{ Success bool }{true})

}


func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")


	session.Values["authenticated"] = false
	session.Values["User"] = nil
	session.Save(r, w)

}


func secretPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/templates/secretPgae.html"))
	session, _ := store.Get(r, "cookie-name")


	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}


	tmpl.Execute(w, nil)


}





func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/register", Register)
	http.HandleFunc("/secretPage", secretPage)
	http.HandleFunc("/logout", logout)

	log.Print("Listening....")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}