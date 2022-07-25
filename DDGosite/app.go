package main

import (
	"database/sql"
	"log"
	"fmt"
	"strings"
	"html/template"
	"net/http"
	"time"
	"golang.org/x/crypto/bcrypt"
	"github.com/gorilla/sessions"
	_ "github.com/go-sql-driver/mysql"

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

func QueryHandler(query string) {
	// testing for bugs
	res1 := strings.Contains(query, "=")
    res2 := strings.Contains(query, "-")
    res3 := strings.Contains(query, ";")
    res4 := strings.Contains(query, ":")
    res5 := strings.Contains(query, "'")

    if res1 == true{
    	return false
    }else if res2 == true{
    	return false
    }else if res3 == true{
    	return false
    }else if res4 == true{
    	return false
    } else if res5 == true{
    	return false
    }else {
    	return true
    }

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

	db, err := sql.Open("mysql", "Test:toor@(127.0.0.1:3308)/?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("USE demon")
	if err != nil {
	log.Fatal(err)
	}

	data := LoginRequest{
		Username: r.FormValue("UserName"),
		Password: r.FormValue("PassWord"),

	}

	var (
		
		UN string
		PW string
		
	)
	
	_ = data

	fmt.Println(data)

	// need to test for bugs

	strcheck := QueryHandler(data.Username)

	if strcheck == false{
		fmt.Println("string Contains illegal characters")

		//inform user that those chars are illegal

		tmpl.Execute(w, nil)

	}

	// Cleanput := strings.Replace(data.Username)


	userCheck, _ := db.Query(`SELECT username, password FROM users WHERE username = ?`, data.Username)

	defer userCheck.Close()


	for userCheck.Next() {
		err := userCheck.Scan(&UN, &PW)
		if err != nil {
			log.Fatal(err)
		}

	}

	// check if password hash matches database 


	// user := userCheck

	fmt.Println(UN)
	fmt.Println(PW)

	if UN == data.Username {

		Match := CheckPasswordHash(data.Password, PW)

		if Match == true {

			session.Values["authenticated"] = true
			session.Values["User"] = data.Username
			session.Save(r, w)

			// redirect user to secretpage

			tmpl := template.Must(template.ParseFiles("static/Templates/secretPage.html"))
			tmpl.Execute(w, nil)
			return

		} else if Match == false {
			// flash error message
			fmt.Println("Invalid Password")
		}
	}else {
		// flash error message
		fmt.Println("Invalid Username")
	}

	// tmpl.Execute(w, nil)

	tmpl.Execute(w, struct{ Success bool }{true})
}

func Register(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/Templates/register.html"))

	session, _ := store.Get(r, "cookie-name")

	db, err := sql.Open("mysql", "Test:toor@(127.0.0.1:3308)/?parseTime=true")

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("USE demon")
	if err != nil {
		log.Fatal(err)
	}

	data := RegisterDetails{
		Email: r.FormValue("email"),
		Username: r.FormValue("UserName"),
		Password: r.FormValue("PassWord"),
		ConfPass: r.FormValue("pass2"),
	}


	var (
		UN string
		PW string
		EM string

	)


	_ = data

	fmt.Println(data)


	userCheck, _ := db.Query(`SELECT username, password FROM users WHERE username = ?`, data.Username)

	defer userCheck.Close()

	for userCheck.Next() {
		err := userCheck.Scan(&UN, &PW)
		if err != nil {
			log.Fatal(err)
		}

	}



	// if username or email dose NOT exist in database
	// Hash the password 


	// if statment auths user to grant access to secret page for that session
	// must clean userinput before query
	// this is done to prevent SQL-injections

	if UN == ""{
		if EM == "" {
			Hpass, _ := HashPassword(data.Password)

			result, err := db.Exec(`INSERT INTO users (username, password, email, created_at) VALUES (?, ?, ?, ?)`, data.Username, Hpass, data.Email, time.Now())
			if err != nil {
				log.Fatal(err)
			}
			id, err := result.LastInsertId()
			fmt.Println(id)


			session.Values["authenticated"] = true
			session.Values["User"] = data.Username
			session.Save(r, w)

			tmpl := template.Must(template.ParseFiles("static/Templates/secretPage.html"))
			tmpl.Execute(w, nil)
			return


		}else if EM != "" {
			// flash email is not available
			fmt.Println("Email in already in use")

		}
	}else if UN != "" {
		// flash username  is not available
		fmt.Println("Username is not available")


	}

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


func ToolsPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/templates/tools.html"))
	// session, _ := store.Get(r, "cookie-name")




	tmpl.Execute(w, nil)

}

func DAOPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/templates/DemonDAO.html"))

	tmpl.Execute(w, nil)
}

func Projects(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/templates/Projects.html"))

	tmpl.Execute(w, nil)
}


func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/register", Register)

	http.HandleFunc("/logout", logout)
	http.HandleFunc("/secretPage", secretPage)
	http.HandleFunc("/ToolsPage", ToolsPage)
	http.HandleFunc("/DemonDAO", DAOPage)
	http.HandleFunc("/Projects", Projects)

	log.Print("Listening....")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}