package main 

import (
	"database/sql"
	"fmt"
	"strings"
	"net/http"
	"html/template"
	"log"
	"time"
	// "reflect"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/go-sql-driver/mysql"
)

var tpl *template.Template

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


// use cach/database
var sessions = map[string]Session{}


// session Values
type Session struct {
	Authenticated bool
	username string
	expiry time.Time 
}



func (s Session) Expired() bool {
	return s.expiry.Before(time.Now())
}



func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 16)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}



func login(w http.ResponseWriter, r *http.Request) {

	// connect to Database
	db, err := sql.Open("mysql", "Test:toor@(127.0.0.1:3308)/?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("USE goauth")
	if err != nil {
	log.Fatal(err)
	}

	data := LoginRequest{
		Username: r.FormValue("UserName"),
		Password: r.FormValue("PassWord"),

	}
	_ = data

	UNlower := strings.ToLower(data.Username)

	var (
		
		UN string
		PW string
		
	)


	// issues with database
	userCheck, _ := db.Query(`SELECT username, password FROM users WHERE username = ?`, UNlower)

	defer userCheck.Close()

	for userCheck.Next() {
		err := userCheck.Scan(&UN, &PW)
		if err != nil {
			log.Fatal(err)
		}

	}


	if UN == UNlower {

		if Match := CheckPasswordHash(data.Password, PW); Match == true{

			seshToken := uuid.NewString()
			expiresAt := time.Now().Add(120 * time.Second)


			// add seshToken to session cach/database
			// since im not using cach/database
			// mapping to sessions map 
			sessions[seshToken] = Session{
				Authenticated: true,
				username: UNlower,
				expiry: expiresAt,
			}


			//set cookie
			http.SetCookie(w, &http.Cookie{
				Name: "Session_token",
				Value: seshToken,
				Expires: expiresAt,
			})

			tpl.ExecuteTemplate(w, "Home.html", "null")
			return

		}

		fmt.Println("Invalid Password")
		tpl.ExecuteTemplate(w, "login.html", "null")
		return
	} else if UN != UNlower {

		fmt.Println("Invalid Username")
		tpl.ExecuteTemplate(w, "login.html", "null")
		return
	}

	tpl.ExecuteTemplate(w, "login.html", "null")


}

func logout(w http.ResponseWriter, r *http.Request) {
	cook, err := r.Cookie("Session_token")
	if err != nil {
		if err == http.ErrNoCookie {

			tpl.ExecuteTemplate(w, "index.html", "null")

			return
		}

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	seshToken := cook.Value

	delete(sessions, seshToken)

	http.SetCookie(w, &http.Cookie{
		Name: "Session_token",
		Value: "",
		Expires: time.Now(),
	})

	tpl.ExecuteTemplate(w, "index.html", "null")

}

func Register(w http.ResponseWriter, r *http.Request) {

	// connect to Database
	db, err := sql.Open("mysql", "Test:toor@(127.0.0.1:3308)/?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("USE goauth")
	if err != nil {
	log.Fatal(err)
	}


	data := RegisterDetails{
		Email: r.FormValue("email"),
		Username: r.FormValue("UserName"),
		Password: r.FormValue("PassWord"),
		ConfPass: r.FormValue("pass2"),
	}


	_ = data

	UNlower := strings.ToLower(data.Username)

	var (
		UN string
		PW string
		EM string
	)



	userCheck, _ := db.Query(`SELECT username, password FROM users WHERE username = ?`, UNlower)
	defer userCheck.Close()

	// scaning values from Database
	for userCheck.Next() {
		err := userCheck.Scan(&UN, &PW)
		if err != nil {
			log.Fatal(err)
		}

	}

	EmailCheck, _ := db.Query(`SELECT email FROM users WHERE email = ?`, data.Email)
	defer EmailCheck.Close()

	for EmailCheck.Next() {
		err := EmailCheck.Scan(&EM)
		if err != nil {
			log.Fatal(err)
		}
	}

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

							seshToken := uuid.NewString()
							expiresAt := time.Now().Add(120 * time.Second)


							sessions[seshToken] = Session{
								Authenticated: true,
								username: UNlower,
								expiry: expiresAt,
							}


							//set cookie
							http.SetCookie(w, &http.Cookie{
								Name: "Session_token",
								Value: seshToken,
								Expires: expiresAt,
							})

							tpl.ExecuteTemplate(w, "Home.html", "auth")
							return



						}else if EM != "" {

							tpl.ExecuteTemplate(w, "register.html", "null")
							fmt.Println("Email in already in use")
							return

						}
					}else if UN != "" {
						tpl.ExecuteTemplate(w, "register.html", "null")
						fmt.Println("Username is not available")
						return
					}

				}else if data.Email == ""{
					tpl.ExecuteTemplate(w, "register.html", "null")
					fmt.Println("Username is not available")
					return
				}
			}else{
				tpl.ExecuteTemplate(w, "register.html", "null")
				fmt.Println("Passwords are not the same")
				return
			}

		}else if data.Password == "" {
			tpl.ExecuteTemplate(w, "register.html", "null")
			fmt.Println("Password is empty")
			return
		}
	}else if data.Username == "" {
		tpl.ExecuteTemplate(w, "register.html", "null")
		fmt.Println("Username is empty")
		return
	}

	tpl.ExecuteTemplate(w, "register.html", "mull")



}

func Home(w http.ResponseWriter, r *http.Request) {

	cook, err := r.Cookie("Session_token")
	if err != nil {
		if err == http.ErrNoCookie {

			tpl.ExecuteTemplate(w, "login.html", "null")
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	seshToken := cook.Value

	Usesh, exists := sessions[seshToken]
	if !exists {

		tpl.ExecuteTemplate(w, "index.html", "null")
		return
	}

	if Usesh.Expired(){
		delete(sessions, seshToken)

		tpl.ExecuteTemplate(w, "index.html", "null")
		return
	}


	tpl.ExecuteTemplate(w, "Home.html", "auth")


}

func SeshRefresh(w http.ResponseWriter, r *http.Request) {

	cook, err := r.Cookie("Session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	seshToken := cook.Value

	Usesh, exists := sessions[seshToken]
	if !exists {
		http.Redirect(w, r, "/login", http.StatusFound)
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

func IDX(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", "null")
}


func main() {

	tpl, _ = template.ParseGlob("./static/Templates/*html")

	http.HandleFunc("/", IDX)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/Home", Home)
	http.HandleFunc("/refresh", SeshRefresh)
	http.HandleFunc("/register", Register)

	log.Print("Listening....")
	log.Fatal(http.ListenAndServe(":8080", nil))


}