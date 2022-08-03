package main

import (
	"database/sql"
	"log"
	"fmt"
	"strings"
	// "crypto/sha256"
	// "sync"
	// "container/list"
	"html/template"
	"net/mail"
	"net/http"
	"time"
	"golang.org/x/crypto/bcrypt"
	"github.com/gorilla/sessions"
	_ "github.com/go-sql-driver/mysql"

)

var tpl *template.Template 

var (
	key = []byte("923ef6d931b1d39b9db72a224567f080f3360c493ac72295378fcb43e8cda6c2")
	store = sessions.NewCookieStore(key)
	// user = nil
	// store = sessions.NewCookieStore(user)
)

var store2 = sessions.NewCookieStore(
	[]byte("new-Auth-key"),
	[]byte("new-encryption-key"),
	[]byte("old-auth-key"),
	[]byte("old-encryption-key"),
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

func validateEmail(addy string) (strin, bool) {
	addr, err := mail.ParseAddress(addy)
	if err != nil {
		return "", false

	}
	return addr.Address, true
}


func QueryHandler(query string) bool{
	// testing for bugs

	// var Ichar = []string{"=","-",";",":","'"}

	// var i = 0
	// for i < len(Ichar){
	// 	res1 := strings.Contains(query, Ichar[i])
	// 	if res1 == true{
	// 		return false
	// 	}else{
	// 		return true
	// 	}
	// 	i++
	// }
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
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 16)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}





func loginPage(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("static/Templates/Login.html"))

	// session, _ := r.Cookie("cookie-name")

	//add err

	// sessionToken := "generate UI for sessions"



	// connect to Database
	db, err := sql.Open("goauth", "Test:toor@(127.0.0.1:3308)/?parseTime=true")
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

	//
	// tmpl.Execute(w, struct{ Success bool }{false})

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
	// test user input for "=-;:'"
	// maybe temporary 
	strcheck := QueryHandler(data.Username)

	if strcheck == false{
		// flash invalid use of chars "=-;:'"
		fmt.Println("string Contains illegal characters")

		//inform user that those chars are illegal

		tmpl.Execute(w, struct{ Success bool }{true})
	}


	UNlower := strings.ToLower(data.Username)
	// Cleanput := strings.Replace(data.Username)


	// Query user from Database
	userCheck, _ := db.Query(`SELECT username, password FROM users WHERE username = ?`, UNlower)

	defer userCheck.Close()

	for userCheck.Next() {
		err := userCheck.Scan(&UN, &PW)
		if err != nil {
			log.Fatal(err)
		}

	}



	fmt.Println(UN)
	fmt.Println(PW)


	// if UN was in database this statment would pass
	if UN == UNlower {
	
		// checks if user input the password correctly
		Match := CheckPasswordHash(data.Password, PW)


		// auths the session and routs user to secret page
		if Match == true {

			session, _ := store.Get(r, "cookie-name")

			// redirect user to secretpage
			session.Values["authenticated"] = true
			session.Values["User"] = UNlower
			session.Save(r, w)

			tpl.ExecuteTemplate(w, "secretPage.html", "auth")


			// http.Redirect(w, r, "/secretPage", http.StatusFound)

			return

		} else if Match == false {
			// flash error message
			// session.AddFlash("Invalid Password")
			fmt.Println("Invalid Password")
			tmpl.Execute(w, struct{ Success bool }{true})
		}
	}else {
		// flash error message
		// session.AddFlash("Invalid Username")
		fmt.Println("Invalid Username")
		tmpl.Execute(w, struct{ Success bool }{true})
	}

	// tmpl.Execute(w, struct{ Success bool }{false})

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

	// -----------------------------
	// tmpl.Execute(w, struct{ Success bool }{false})
	// -----------------------------


	// input from front end form
	data := RegisterDetails{
		Email: r.FormValue("email"),
		Username: r.FormValue("UserName"),
		Password: r.FormValue("PassWord"),
		ConfPass: r.FormValue("pass2"),
	}

	// Variables to store Database values

	var (
		UN string
		PW string
		EM string

	)


	_ = data

	fmt.Println(data)


	strcheck := QueryHandler(data.Username)

	if strcheck == false{
		// flash invalid use of chars "=-;:'"
		fmt.Println("string Contains illegal characters")

		//inform user that those chars are illegal

		tmpl.Execute(w, struct{ Success bool }{true})
	}


	UNlower := strings.ToLower(data.Username)

	// Query database to see if username exist

	userCheck, _ := db.Query(`SELECT username, password FROM users WHERE username = ?`, UNlower)

	defer userCheck.Close()

	// scaning values from Database
	for userCheck.Next() {
		err := userCheck.Scan(&UN, &PW)
		if err != nil {
			log.Fatal(err)
		}

	}

	// could probally join query

	EmailCheck, _ := db.Query(`SELECT email FROM users WHERE email = ?`, data.Email)
	
	// Vmail := validateEmail(EmailCheck).(bool)
	// if Vmail 

	defer EmailCheck.Close()

	for EmailCheck.Next() {
		err := EmailCheck.Scan(&EM)
		if err != nil {
			log.Fatal(err)
		}
	}


	// I can clean this up using some frontend code 

	// if any values from the form are empty then the page informs the user
	if data.Username != ""{
		if data.Password != ""{
			if data.ConfPass == data.Password {
				if data.Email != "" {

					// if Username and Email do not exist in Database
					// then the program hashes and stores the user and routs them to secret page
					if UN == ""{

						if EM == "" {


							Hpass, _ := HashPassword(data.Password)

							result, err := db.Exec(`INSERT INTO users (username, password, email, created_at) VALUES (?, ?, ?, ?)`, UNlower, Hpass, data.Email, time.Now())
							if err != nil {
								log.Fatal(err)
							}
							id, err := result.LastInsertId()
							fmt.Println(id)


							session.Values["authenticated"] = true
							session.Values["User"] = UNlower
							session.Save(r, w)

							http.Redirect(w, r, "static/Templates/secretPage.html", http.StatusFound)
							// tmpl := template.Must(template.ParseFiles("static/Templates/secretPage.html"))
							// tmpl.Execute(w, session)

							return


						}else if EM != "" {
							// flash email is not available
							session.AddFlash("Email is not available")
							tmpl.Execute(w, nil)
							fmt.Println("Email in already in use")

						}
					}else if UN != "" {
						// flash username  is not available
						session.AddFlash("Username is not available")
						tmpl.Execute(w, nil)
						fmt.Println("Username is not available")
					}

				}else if data.Email == ""{
					// flash email is empty
					session.AddFlash("Email can not be empty")
					tmpl.Execute(w, struct{ Success bool }{false})
					fmt.Println("Username is not available")
				}
			}else{
				//flash password is not the same
				session.AddFlash("password is not the same")
				tmpl.Execute(w, struct{ Success bool }{false})
				fmt.Println("Passwords are not the same")
			}

		}else if data.Password == "" {
			// flash password is empty
			session.AddFlash("password can not be empty")
			tmpl.Execute(w, struct{ Success bool }{false})
			fmt.Println("Password is empty")
		}
	}else if data.Username == "" {
		// flasj Username is empty
		session.AddFlash("Username can not be empty")
		tmpl.Execute(w, struct{ Success bool }{false})
		fmt.Println("Username is empty")
	}
	
	// tmpl.Execute(w, struct{ Success bool }{false})

}


func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")


	session.Values["authenticated"] = false
	session.Values["User"] = nil
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusFound)

}


func secretPage(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	fmt.Println(session)



	test := session.Values["user"]
	if test == nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} 
	// else if test is a registered user then 

	tmpl := template.Must(template.ParseFiles("static/templates/tools.html"))

	tmpl.Execute(w, nil)

}


func ToolsPage(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	tmpl := template.Must(template.ParseFiles("static/templates/tools.html"))
	// session, _ := store.Get(r, "cookie-name")

	if session.Values["authenticated"] == true {
		fmt.Println(session.Values["user"])
	}


	tmpl.Execute(w, nil)

}

func DAOPage(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	tmpl := template.Must(template.ParseFiles("static/templates/DemonDAO.html"))


	if session.Values["authenticated"] == true {
		fmt.Println(session.Values["user"])
	}

	
	tmpl.Execute(w, nil)
}

func Projects(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	tmpl := template.Must(template.ParseFiles("static/templates/Projects.html"))

	
	if session.Values["authenticated"] == true {
		fmt.Println(session.Values["user"])
	}


	tmpl.Execute(w, nil)
}


func main() {

	tpl, _ = template.ParseGlob("./static/Templates/*html")

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/register", Register)

	http.HandleFunc("/logout", logout)
	http.HandleFunc("/SecretPage", secretPage)
	http.HandleFunc("/ToolsPage", ToolsPage)
	http.HandleFunc("/DemonDAO", DAOPage)
	http.HandleFunc("/Projects", Projects)

	log.Print("Listening....")



	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}