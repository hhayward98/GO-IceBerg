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

func SeshRefresh(w http.ResponseWriter, r *http.Request) {
	Basic := HTMLDATASET()
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

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 16)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
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
	if len(FormData.Username) == 0 {
		if len(FormData.Password) == 0 {
			// loading page should hit here
			tpl.ExecuteTemplate(w, "login.html", Basic)
			return
		}else if (FormData.Password) > 0 {
			log.Print("username is empty")
			Basic.Body = "Username can not be empty!"
			tpl.ExecuteTemplate(w, "login.html", Basic)
			return
		}
	}else if len(FormData.Username) > 0 {
		if len(FormData.Password) == 0 {
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

		if Match := CheckPasswordHash(FormData.Password, PW); Match == true{

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
		return
	}

	if QueryHandler(UNlower) != true {
		log.Print("invalid characters detected!!")
		Basic.Body = "Illegal characters detected!!"
		tpl.ExecuteTemplate(w, "login.html", Basic)
		return
	}


	// checking form data 
	if len(FormData.Email) == 0 {
		if len(FormData.Username) == 0 {
			if len(FormData.Password) == 0 {
				if len(FormData.ConfPass) == 0 {
					// 
					tpl.ExecuteTemplate(w, "register.html", Basic)
					return
				}else {
					log.Print("Information missing!")
					Basic.Body = "missing Information!"
					tpl.ExecuteTemplate(w, "register.html", Basic)
					return
				}
			}
		}
		// While Email is empty 
		log.Print("Email is empty")
		Basic.Body = "Email can not be empty!"
		tpl.ExecuteTemplate(w, "register.html", Basic)
		return
	}else if len(FormData.Email) > 1 {
		if len(FormData.Username) > 1 {
			if len(FormData.Password) > 1 {
				if FormData.Password != FormData.ConfPass {
					log.Print("Passwords do not match!")
					Basic.Body = "Passwords do not match!"
					tpl.ExecuteTemplate(w, "register.html", Basic)
					return
				}
			}else if len(FormData.Password) == 0 {
				log.Print("Password is Empty")
				Basic.Body = "Password can not be empty"
				tpl.ExecuteTemplate(w, "register.html", Basic)
				return
			}

		}else if len(FormData.Username) == 0 {
			log.Print("Username is Empty")
			Basic.Body = "Username Can not be empty!"
			tpl.ExecuteTemplate(w, "register.html", Basic)
			return
		}
	}
	// continues if form data is good


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



func logout(w http.ResponseWriter, r *http.Request) {
	log.Print("Loging user out...")
	Basic := HTMLDATASET()
	cook, err := r.Cookie("Session_token")
	if err != nil {
		if err == http.ErrNoCookie {

			tpl.ExecuteTemplate(w, "index.html", Basic)
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

	tpl.ExecuteTemplate(w, "index.html", Basic)
	return
}



func ToolsPage(w http.ResponseWriter, r *http.Request) {
	log.Print("Running ToolsPage...")

	Basic := HTMLDATASET()

	cook, err := r.Cookie("Session_token")
	if err != nil {
		if err == http.ErrNoCookie {

			tpl.ExecuteTemplate(w, "tools.html", Basic)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	seshToken := cook.Value

	Usesh, exists := sessions[seshToken]
	if !exists {

		tpl.ExecuteTemplate(w, "tools.html", Basic)
		return
	}

	if Usesh.Expired(){
		delete(sessions, seshToken)
	}
	tpl.ExecuteTemplate(w, "tools.html", Basic)
	return
}


func DAOPage(w http.ResponseWriter, r *http.Request) {
	log.Print("Running DAOPage...")
	Basic := HTMLDATASET()
	cook, err := r.Cookie("Session_token")
	if err != nil {
		if err == http.ErrNoCookie {

			tpl.ExecuteTemplate(w, "DemonDAO.html", Basic)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	seshToken := cook.Value

	Usesh, exists := sessions[seshToken]
	if !exists {

		tpl.ExecuteTemplate(w, "DemonDAO.html", Basic)
		return
	}

	if Usesh.Expired(){
		delete(sessions, seshToken)

		tpl.ExecuteTemplate(w, "DemonDAO.html", Basic)
		return
	}

	tpl.ExecuteTemplate(w, "DemonDAO.html", Basic)
	return
}


func Home(w http.ResponseWriter, r *http.Request) {
	log.Print("Running Home Page...")

	Basic := HTMLDATASET()

	cook, err := r.Cookie("Session_token")
	if err != nil {
		if err == http.ErrNoCookie {

			tpl.ExecuteTemplate(w, "index.html", Basic)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	seshToken := cook.Value

	Usesh, exists := sessions[seshToken]
	if !exists {

		tpl.ExecuteTemplate(w, "index.html", Basic)
		return
	}

	if Usesh.Expired(){
		delete(sessions, seshToken)

		tpl.ExecuteTemplate(w, "index.html", Basic)
		return
	}

	tpl.ExecuteTemplate(w, "index.html", Basic)
	return

}

func Projects(w http.ResponseWriter, r *http.Request) {
	log.Print("Running Projects Page...")

	Basic := HTMLDATASET()

	cook, err := r.Cookie("Session_token")
	if err != nil {
		if err == http.ErrNoCookie {

			tpl.ExecuteTemplate(w, "Projects.html", Basic)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	seshToken := cook.Value

	Usesh, exists := sessions[seshToken]
	if !exists {

		tpl.ExecuteTemplate(w, "Projects.html", Basic)
		return
	}

	if Usesh.Expired(){
		delete(sessions, seshToken)

		tpl.ExecuteTemplate(w, "Projects.html", Basic)
		return
	}
	
	tpl.ExecuteTemplate(w, "Projects.html", Basic)
	return

}


func AppRoutes() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", Home)
	http.HandleFunc("/Login", Login)
	http.HandleFunc("/Register", Register)
	http.HandleFunc("/Logout", logout)
	http.HandleFunc("/Tools", ToolsPage)
	http.HandleFunc("/Projects", Projects)
	http.HandleFunc("/DemonDAO", DAOPage)


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