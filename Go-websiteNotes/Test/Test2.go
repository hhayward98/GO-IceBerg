package main


import (
    "database/sql"
    "fmt"
    "log"
    "time"
	"golang.org/x/crypto/bcrypt"

    _ "github.com/go-sql-driver/mysql"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 32)
	return string(bytes), err
}

func main() {
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


	var (
		Username string
		Password string
		Email string
	)
	fmt.Println("_________________")
	fmt.Println("User Register Sim")
	fmt.Println("_________________\n")
	fmt.Println("Username: ")
	fmt.Scanln(&Username)
	fmt.Println("Password")
	fmt.Scanln(&Password)
	fmt.Println("Email")
	fmt.Scanln(&Email)

	var (
		UN string
		PW string
		EM string

	)



	userCheck, _ := db.Query(`SELECT username, password, Email FROM users WHERE username = ?`, Username)

	defer userCheck.Close()

	for userCheck.Next() {
		err := userCheck.Scan(&UN, &PW, &EM)
		if err != nil {
			log.Fatal(err)
		}

	}


	if Username != ""{
		if Password != ""{
			if Email != ""{
				if UN == ""{
					if EM == "" {
						Hpass, _ := HashPassword(Password)

						result, err := db.Exec(`INSERT INTO users (username, password, email, created_at) VALUES (?, ?, ?, ?)`, Username, Hpass, Email, time.Now())
						if err != nil {
							log.Fatal(err)
						}
						id, err := result.LastInsertId()
						fmt.Println(id)


						fmt.Println("User registration complete")


					}else if EM != "" {
						// flash email is not available
						fmt.Println("Email in already in use")

					}
				}else if UN != "" {
					// flash username  is not available
					fmt.Println("Username is not available")


				}


			}else if Email == "" {
				fmt.Println("Email can not be Empty")
			}
		}else if Password == "" {
			fmt.Println("Password can not be Empty")
		}
	}

	fmt.Println(UN)
	fmt.Println(EM)
	fmt.Println(PW)


}