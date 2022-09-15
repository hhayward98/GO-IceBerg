package LoginC

import (
	"database/sql"
	"fmt"
	"strings"
	"log"

	"golang.org/x/crypto/bcrypt"

	_ "github.com/go-sql-driver/mysql"
)


type LoginRequest struct {
	Username string
	Password string
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 16)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}


func DataBaseCheck(Uname string, Pword string) {

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
		
	)

	userCheck, _ := db.Query(`SELECT username, password FROM users WHERE username = ?`, Uname)


	defer userCheck.Close()



	for userCheck.Next() {
		err := userCheck.Scan(&UN, &PW)
		if err != nil {
			log.Fatal(err)
		}

	}

	if Uname == UN {

		if Match := CheckPasswordHash(data.Password, PW); Match == true{

			log.Print("Login Success!")
			return "Login Success!"
		}
		log.Print("Invalid Password")
		return "Invalid Password"
	}else if UN != UNlower {
		log.Print("Invalid Username")
		return "Invalid Username"
	}

	return ""
}


func UserInput(Uname string, Pword string) {
	UNlower := strings.ToLower(Uname)

	hpass, err := HashPassword(Pword)
	if err != nil {
		log.Fatal(err)
	}


	status := DataBaseCheck(UNlower, hpass)

	return status

}