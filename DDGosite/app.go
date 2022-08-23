package main

import (
	"database/sql"
	"log"
	"fmt"
	"strings"
	// "regexp"
	"html/template"
	"net/mail"
	"net/http"
	"time"
	"golang.org/x/crypto/bcrypt"
	"github.com/google/uuid"
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


type InputError struct {
	email string
	username string
	password string
	ConfPass string
}


func (s Session) Expired() bool {
	return s.expiry.Before(time.Now())
}


func validateEmail(addy string) (string, bool) {
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

	Eput := InputError{"","Invalid Username!!", "Invalid Password!!", "passwords do not Match!!"}

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

			tpl.ExecuteTemplate(w, "secretPage.html", "null")
			return

		}

		if len(data.Username) > 1 {
			fmt.Println("Invalid Password")
			tpl.ExecuteTemplate(w, "Login.html", Eput.password)
			return
		} 

		tpl.ExecuteTemplate(w, "Login.html", "")
		return
	} else if UN != UNlower {

		fmt.Println("Invalid Username")
		tpl.ExecuteTemplate(w, "Login.html", Eput.username)
		return
	}

	tpl.ExecuteTemplate(w, "Login.html", "")


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
	return
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

	// // regular expression for valid email 
 //    if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, r.Form.Get("email")); !m {
 //        fmt.Println("no")
 //    }else{
 //        fmt.Println("yes")
 //    }

 //    // checks if username is empty
 //    if len(r.Form["UserName"][0]) == 0 {
 //    	fmt.Println("Username is empty")
 //    }

 	Eput := InputError{"Email already in use!!","Username already in use", "Password cant be empty", "passwords do not Match!!"}
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

							tpl.ExecuteTemplate(w, "secretPage.html", "auth")
							return



						}else if EM != "" {

							tpl.ExecuteTemplate(w, "register.html", Eput.email)
							fmt.Println("Email in already in use")
							return

						}
					}else if UN != "" {
						tpl.ExecuteTemplate(w, "register.html", Eput.username)
						fmt.Println("Username is not available")
						return
					}

				}else if data.Email == ""{
					tpl.ExecuteTemplate(w, "register.html", Eput.username)
					fmt.Println("Username is not available")
					return
				}
			}else{
				tpl.ExecuteTemplate(w, "register.html", Eput.ConfPass)
				fmt.Println("Passwords are not the same")
				return
			}

		}else if data.Password == "" {
			tpl.ExecuteTemplate(w, "register.html", Eput.password)
			fmt.Println("Password is empty")
			return
		}
	}else if data.Username == "" {
		if len(data.Email) > 1 {
			tpl.ExecuteTemplate(w, "register.html", Eput.username)
			fmt.Println("Username is empty")
			return
		}
		tpl.ExecuteTemplate(w, "register.html", "")
		return
	}

	tpl.ExecuteTemplate(w, "register.html", "")
	return


}

func Home(w http.ResponseWriter, r *http.Request) {

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

	tpl.ExecuteTemplate(w, "index.html", "auth")
	return

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


func secretPage(w http.ResponseWriter, r *http.Request) {
	cook, err := r.Cookie("Session_token")
	if err != nil {
		if err == http.ErrNoCookie {

			tpl.ExecuteTemplate(w, "Login.html", "null")
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	seshToken := cook.Value

	Usesh, exists := sessions[seshToken]
	if !exists {

		tpl.ExecuteTemplate(w, "Login.html", "null")
		return
	}

	if Usesh.Expired(){
		delete(sessions, seshToken)

		tpl.ExecuteTemplate(w, "Login.html", "null")
		return
	}

	tpl.ExecuteTemplate(w, "secretPage.html", "auth")
	return

}


func ToolsPage(w http.ResponseWriter, r *http.Request) {
	cook, err := r.Cookie("Session_token")
	if err != nil {
		if err == http.ErrNoCookie {

			tpl.ExecuteTemplate(w, "tools.html", "null")
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	seshToken := cook.Value

	Usesh, exists := sessions[seshToken]
	if !exists {

		tpl.ExecuteTemplate(w, "tools.html", "null")
		return
	}

	if Usesh.Expired(){
		delete(sessions, seshToken)
	}
	tpl.ExecuteTemplate(w, "tools.html", "null")
	return
}

func DAOPage(w http.ResponseWriter, r *http.Request) {
	cook, err := r.Cookie("Session_token")
	if err != nil {
		if err == http.ErrNoCookie {

			tpl.ExecuteTemplate(w, "DemonDAO.html", "null")
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	seshToken := cook.Value

	Usesh, exists := sessions[seshToken]
	if !exists {

		tpl.ExecuteTemplate(w, "DemonDAO.html", "null")
		return
	}

	if Usesh.Expired(){
		delete(sessions, seshToken)

		tpl.ExecuteTemplate(w, "DemonDAO.html", "null")
		return
	}

	tpl.ExecuteTemplate(w, "DemonDAO.html", "null")
	return
}




func Projects(w http.ResponseWriter, r *http.Request) {
	cook, err := r.Cookie("Session_token")
	if err != nil {
		if err == http.ErrNoCookie {

			tpl.ExecuteTemplate(w, "Projects.html", "null")
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	seshToken := cook.Value

	Usesh, exists := sessions[seshToken]
	if !exists {

		tpl.ExecuteTemplate(w, "Projects.html", "null")
		return
	}

	if Usesh.Expired(){
		delete(sessions, seshToken)

		tpl.ExecuteTemplate(w, "Projects.html", "null")
		return
	}
	
	tpl.ExecuteTemplate(w, "Projects.html", "null")
}

func members(w http.ResponseWriter, r *http.Request) {
	cook, err := r.Cookie("Session_token")
	if err != nil {
		if err == http.ErrNoCookie {

			tpl.ExecuteTemplate(w, "members.html", "null")
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	seshToken := cook.Value

	Usesh, exists := sessions[seshToken]
	if !exists {

		tpl.ExecuteTemplate(w, "members.html", "null")
		return
	}

	if Usesh.Expired(){
		delete(sessions, seshToken)

		tpl.ExecuteTemplate(w, "members.html", "null")
		return
	}
	
	tpl.ExecuteTemplate(w, "members.html", "null")
}



func main() {

	tpl, _ = template.ParseGlob("./static/Templates/*html")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", Home)

	http.HandleFunc("/login", login)
	http.HandleFunc("/register", Register)

	http.HandleFunc("/logout", logout)
	http.HandleFunc("/SecretPage", secretPage)
	http.HandleFunc("/ToolsPage", ToolsPage)
	http.HandleFunc("/DemonDAO", DAOPage)
	http.HandleFunc("/Projects", Projects)

	log.Print("Listening....")



	log.Print("Listening....")
	err := http.ListenAndServeTLS(":9000", "localhost.crt", "localhost.key", nil)
	if err != nil {
			log.Fatal("ListenAndServe: ", err)
	}


}


// docker run -p 8080:8080 -it dockerdd 