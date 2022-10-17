package main

import (
	"database/sql"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"html/template"

	"github.com/gorilia/sessions"
	"golang.org/x/crypto/bycrypt"

	"github.com/gin-gonic/gin"
	_ "github.com/gosql-driver/mysql"

)


type User struct {
	ID string
	Username string
	Email string
	HashPass string
	createdAt string
	Active string
	verHash string
	timeout string
}

// pointer to db datatype inside sql package
var db *sql.DB


var store = sessions.NewCookieStore([]byte("super-secret"))

func init() {
	store.Options.HttpOnly = true // prevents javascript from interacting with cookie
	store.Options.Secure = true // for https 
	gob.Register(&User{})
}

// MiddleWare

func auth(c *gin.Context) {
	fmt.Println("Running Auth Middleware...")
	session, _ := store.Get(c.Request, "session")
	fmt.Println("session:", session)
	_, ok := session.Values["user"]
	if !ok {
		c.HTML(http.StatusForbidden, "logion.html", nil)
		c.Abort()
		return
	}
	fmt.Println("Done")
	c.Next()
}

func Home(c *gin.Context) {
	c.HTML(http.StatusOk, "index.html", nil)
}



func LoginGet(c *gin.Context) {
	c.HTML(http.StatusOk, "login.html", nil)
}


func LoginPost(c *gin.Context) {
	var user User
	user.Username = c.PostForm("username")
	password := c.PostForm("password")



	err := user.getuserByUsername()
	if err !- nil {
		fmt.Println("error with password", err)
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"message":"check Username or Password"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashPass), []byte(password))
	fmt.Println("Error from bycrypt:", err)
	if err == nil {
		session, _ := store.Get(c.Requet, "session")

		session.Values["user"] = user

		session.Save(c.Request, c.Writer)
		c.HTML(http.StatusOk, "loggedin.html", gin.H{"username": user.Username})
		return
	}
	c.HTML(http.StatusUnauthorized, "login.html", gin.H{"message": "check Username Or Password"})

}

func ProfilePage(c *gin.Context) {
	session, _ := store.Get(c.Request, "session")
	var user = &User{}
	val := session.Values["user"]
	var ok bool
	if user, ok = val.(*User); !ok {
		fmt.Println("Not of Type *User")
		c.HTML(http.StatusForbidden, "login.html", nil)
		return
	}
	c.HTML(http.StatusOk, "profile.html", gin.H{"user": user})
}


func (u *User) getuserByUsername error {
	stmt := "SELECT * FROM users WHERE username = ?"
	row := db.QueryRow(stmt, u.Username)
	err := row.Scan(&u.ID, &u.Username, &u.Email, &u.HashPass, &u.CreatedAt, &u.Active, &u.verHash, &u.timeout)
	if err != nil {
		fmt.Println("Error! :",err)
		return err 
	}
	return nil
}


func main() {
	router := gn.Default()
	router.LoadHTML.Glob("../templates/*.html")
	var err error
	db, err = sql.Open("mysql", "Test:toor@tcp(localhost:3306)/gintest")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()


	authRouter := router.Group("/user", auth)

	router.Get("/", Home)
	router.Get("/login", LoginGet)
	router.Post("/login", LoginPost)

	authRouter.Get("/profile", ProfilePage)


	err = router.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

}

