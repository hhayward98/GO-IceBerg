package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	_ "github.com/go-sql-driver/mysql"


)

var tpl *template.Template


type LoginDetails struct {
	Username string
	Password string
}

type RegisterDetails struct {
	Email string
	Username string
	Password string
	ConfPass string
}

type HTMLDATA struct {
	Header string
	Body string
	Foot string
}


func Debugger(err error, Etype int) {

	if err != nil {
		// Error type 1 will end the program
		if Etype == 1 {
			log.Fatal(err)
		}else if Etype == 2 { // Error type 2 returns error without killing app
			log.Print("===========================")
			log.Print(err)
			log.Print("===========================")
		}
	}
}


func validateEmail(addy string) (string, bool) {
	addr, err := mail.ParseAddress(addy)
	if err != nil {
		return "", false

	}
	return addr.Address, true
}

func QueryHandler(query string) bool{

	res1 := strings.Contains(query, "=")
    res2 := strings.Contains(query, "-")
    res3 := strings.Contains(query, ";")
    res4 := strings.Contains(query, ":")
    res5 := strings.Contains(query, "'")
    res6 := strings.Contains(query, "`")

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
    }else if res6 == true {
    	return false
    }else {
    	return true
    }

}

func HTMLDATASET() {
	data := HTMLDATA{
		Header: "hostmachineIPaddress",
		Body: "",
		Foot: "",
	}
	_ = data

	return data
}

func Login(w http.ResponseWriter, r *http.Request) {
	log.Print("Running Login Page")

	db, err := sql.Open("mysql", "test:toor@tcp(db:3306)/ddlabs")
	Debugger(err, 1)
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// initiate HTML struct
	Basic := HTMLDATASET()
	fmt.Println(Basic.Header)

	FormData := LoginDetails{
		Username: r.FormValue("UserName")
		Password: r.FormValue("PassWord")
	}
	_ = FormData

	UNlower := strings.ToLower(FormData.Username)
	//
	if QueryHandler(UNlower) != true {
		log.Print("invalid characters detected!!")
		Basic.Body = "Illegal characters detected!!"
		tpl.ExecuteTemplate(w, "login.html", Basic)
		return
	}


	// 
	if len(data.Username) == 0 {
		if len(data.Password) == 0 {
			// loading page should hit here
			tpl.ExecuteTemplate(w, "login.html", Basic)
			return
		}else if (data.Password) > 0 {
			log.Print("username is empty")
			Basic.Body = "Username can not be empty!"
			tpl.ExecuteTemplate(w, "login.html", Basic)
			return
		}
	}else if len(data.Username) > 0 {
		if len(data.Password) == 0 {
			log.Print("Password is empty")
			Basic.Body = "Password can not be empty"
			tpl.ExecuteTemplate(w, "login.html")
			return
		}
	}

	var (
		UN string
		PW string	
	)

	userCheck, _ := db.Query(`SELECT username, password FROM users WHERE username = ?`, UNlower)
	defer userCheck.Close()

	for userCheck.Next() {
		err := userCheck.Scan(&UN, &PW)
		Debugger(err, 1)
	}

	if UN == UNlower {

		if Match := CheckPasswordHash(data.Password, PW); Match == true{

			// Sessions stuff changing next
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
			log.Print("Login Success")
			tpl.ExecuteTemplate(w, "secretPage.html", Basic)
			return


		}
	}else if UN != UNlower {
		log.Print("Invalid Username")
		Basic.Body = "Invalid Username!"
		tpl.ExecuteTemplate(w, "login.html", Basic)
		return
	}

}


func Register(w http.ResponseWriter, r *http.Request) {
	log.Print("Running Register Page...")

	db, err := sql.Open("mysql", "test:toor@tcp(db:3306)/ddlabs")
	Debugger(err, 1)
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// initiate HTML struct
	Basic := HTMLDATASET()

	FormData := RegisterDetails {
		Email: r.FormValue("email")
		Username: r.FormValue("UserName")
		Password: r.FormValue("PassWord")
		ConfPass: r.FormValue("ConfPass")
	}
	_ = FormData

	UNlower := strings.ToLower(FormData.Username)
	addy, BL := validateEmail(FormData.Email)
	if BL == false {
		log.Print("Email is not valid!")
		Basic.Body = "Email is not valid!"
		tpl.ExecuteTemplate(w, "register.html", Basic)
	}

	var (
		UN string
		PW string
		EM string
	)


	// Benchmark Query
	userCheck, _ := db.Query(`SELECT username, password FROM users WHERE username = ?`, UNlower)
	defer userCheck.Close()
	for userCheck.Next() {
		err := userCheck.Scan(&UN, &PW)
		Debugger(err, 1)
	}

	EmailCheck, _ := db.Query(`SELECT email FROM users WHERE email = ?`, data.Email)
	defer EmailCheck.Close()
	for EmailCheck.Next() {
		err := EmailCheck.Scan(&EM)
		Debugger(err, 1)
	}

	// ....
	

}



func ToolsPage(w http.ResponseWriter, r *http.Request) {
	log.Print("Running ToolsPage")

	Basic := HTMLDATASET()

	tpl.ExecuteTemplate(w, "tools.html", Basic)

}


func Home(w http.ResponseWriter, r *http.Request) {
	log.Print("Running Home Page")

	Basic := HTMLDATASET()

	tpl.ExecuteTemplate(w, "index.html", Basic)
	return

}


func AppRoutes() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", Home)
	http.HandleFunc("/Login", Login)
	http.HandleFunc("/Register", Register)

	log.Fatal(http.ListenAndServe(":8080", nil))

	//TLS
	// err := http.ListenAndServeTLS(":9000", "localhost.crt", "localhost.key", nil)
	// Debugger(err, 1)

}


func main() {

	tpl, _ = template.ParseGlob("./static/templates/*html")

	log.Print("Listening....")

	AppRoutes()

}