package main


import (	
	"fmt"
	"log"
	"container/list"
	"github.com/astaxie/session"
	"sync"
	"time"
	_ "github.com/astaxie/session/providers/memory"
)




func login(w http.ResponseWriter, r *http.Request) {


	sesh := globalSessions.SessionStart(w, r)


	if r.Method == "GET" {

		tmpl, _ := template.ParseFiles("/static/secretpage.html")

		tmpl.Execute(w, sesh.Get("username"))		

	} else {

		data := LoginRequest{
			Username: r.FormValue("UserName"),
			Password: r.FormValue("PassWord"),
		}

		sesh.Set("username", r.Form["username"])

		// verify input 
	}

}

func secretPage(w http.ResponseWriter, r *http.Request) {

	sesh := globalSessions.SessionStart(w, r)
	cook := sesh.Get("username")

func init() {
	globalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
	go globalSessions.GC()
}

func Home(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", sesh.Get("username"))
}


func main() {

	var globalSessions *session.SManager

	http.HandleFunc("/", Home)
	http.HandleFunc("/secretpage", secretPage)
	http.HandleFunc("/login", login)

}