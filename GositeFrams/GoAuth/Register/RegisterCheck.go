package RegisterC

import (
	"database/sql"

	"strings"
	"log"
	"net/mail"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/go-sql-driver/mysql"
)


type RegisterDetails struct {
	Email string
	Username string
	Password string
	ConfPass string
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
    } else if res6 == true{
    	return false
    }else {
    	return true
    }

}


func validateEmail(addy string) (string, bool) {
	addr, err := mail.ParseAddress(addy)
	if err != nil {
		return "", false

	}
	return addr.Address, true
}




func InputHandler(Uname string, Email string, Pword string, ConfPass string) string {

// input handling 
	var ErrBuffer string

	UNlower := strings.ToLower(Uname)

	addy, BL := validateEmail(Email)

	if BL == false {
		log.Print(addy)
		ErrBuffer = "Email is not valid!"
		return ErrBuffer

	}

	if QueryHandler(UNlower) != true {
		ErrBuffer = "Illegal characters detected!!"
		return ErrBuffer
	}

	return "Valid"

}


func dbUnameCheck(Uname string, Email string) {

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

	var (
		UN string
		PW string
		EM string
	)


	// Benchmark Query
	userCheck, _ := db.Query(`SELECT username, password FROM users WHERE username = ?`, Uname)
	defer userCheck.Close()
	for userCheck.Next() {
		err := userCheck.Scan(&UN, &PW)
		Debugger(err, 1)
	}

	EmailCheck, _ := db.Query(`SELECT email FROM users WHERE email = ?`, Email)
	defer EmailCheck.Close()
	for EmailCheck.Next() {
		err := EmailCheck.Scan(&EM)
		Debugger(err, 1)
	}

	return

}



func RegisterUser(Uname string, Email string, Pword string, ConfPass string) string {

	var ErrBuffer string

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

	var (
		UN string
		PW string
		EM string
	)


	// Benchmark Query
	userCheck, _ := db.Query(`SELECT username, password FROM users WHERE username = ?`, Uname)
	defer userCheck.Close()
	for userCheck.Next() {
		err := userCheck.Scan(&UN, &PW)
		Debugger(err, 1)
	}

	EmailCheck, _ := db.Query(`SELECT email FROM users WHERE email = ?`, Email)
	defer EmailCheck.Close()
	for EmailCheck.Next() {
		err := EmailCheck.Scan(&EM)
		Debugger(err, 1)
	}

	if Uname != ""{
		if Pword != ""{
			if ConfPass == Pword {
				if Email != "" {

					if UN == "" {

						if EM == "" {


							HashPass, err := HashPassword(Pword)
							Debugger(err, 2)

							result, err := db.Exec(`INSERT INTO users (username, password, email) VALUES (?, ?, ?)`, Uname, HashPass, Email)
							Debugger(err, 1)

							id, err := result.LastInsertId()
							log.Print(id)

							return "Valid"

						}else if EM != ""{

							ErrBuffer = "Email not available"
							return ErrBuffer
						}
					}else if UN != "" {

						ErrBuffer = "Username is not available"
						return ErrBuffer
					}

				}else if Email == ""{

					ErrBuffer = "Email is empty"
					return ErrBuffer
				}
			}else{

				ErrBuffer = "Passwords do not match!"
				return ErrBuffer
			}

		}else if Pword == "" {

			ErrBuffer = "Password is Empty"
			return ErrBuffer
		}
	}else if Uname == "" {
		if len(Email) > 1 {

			ErrBuffer = "Username is Empty"
			return ErrBuffer
		}
		return "null"
	}
	return "null"

}


