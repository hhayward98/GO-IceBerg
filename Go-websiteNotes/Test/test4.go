package main

import (
	"net/http"
	"github.com/gorilla/sessions"
	"fmt"
	"sync"
	"html/template"

)


var store = sessions.NewCookieStore(
	[]byte("new-Auth-key"),
	[]byte("new-encryption-key"),
	[]byte("old-auth-key"),
	[]byte("old-encryption-key"),
)


// Session manager

type LoginRequest struct {
	Username string
	Password string

}

type SManager struct {
	cookieName	string // private 
	lock		sync.Mutex // protects session
	provider	Provider
	maxlifetime int64
}

type Provider interface {
	SessionInit(sid string) (Session, error) // implements the init of a session and returns a new session if success
	SessionRead(sid string) (Session, error) // return a session. creates a new session and returns it if it does not exist
	SessionDestroy(sid string) error // given an sid , deletes the corresponding session
	SessionGC(maxLifeTime int64) // SessionGC deletes expired session variables 
}

type Session interface {
	Set(key, value interface{}) error //set session value
	Get(key interface{}) interface{} //get session value
	Delete(key interface{}) error // delete session value
	sessionId() string	
}

func NewManager(Name, cookieName string, maxlifetime int64) (*SManager, error) {
	provider, ok := provides[Name]
	if !ok {
		return nil, fmt.Errorf("Session: unkown provide %q (forgotten import?)", Name)
	}
	return &SManager{provider: provider, cookieName: cookieName, maxlifetime: maxlifetime}, nil
}

func init() {

	globalSessions = NewManager("memory", "gosessionid", 3600)
}

var provides = make(map[string]Provider)

func (manager *SManager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()

	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		sid := manager.sessionId()
		session, _ =manager.provider.SessionInit(sid)
		cookie := http.Cookie{Name: manager.cookieName, value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(manager.maxlifetime)}
		http.SetCookie(w, &cookie)
	}else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.provider.SessionRead(sid)
	}
	return
}


func secretPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/templates/secretPgae.html"))

	fmt.Println(Session.sessionId())



	tmpl.Execute(w, nil)


}

func login(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("forms.html"))

	sesh := globalSessions.SessionStart(w,r)

	data := LoginRequest{
		Username: r.FormValue("Username"),
		Password: r.FormValue("Password"),
	}

	_ = data

	if data.Username != "" {
		if data.Password != ""{
			sesh.Set("authenticated", true)
			sesh.Set("user", data.Username)

			tmpl.Execute(w, sesh)
			fmt.Println(data.Username)
			fmt.Println(data.Password)
		}else {
			fmt.Println("Password cant be null")
			tmpl.Execute(w, struct{ Success bool }{false})
		}
	}else {
		fmt.Println("Username cant be null")
		tmpl.Execute(w, struct{ Success bool }{false})
	}

}

func main() {
	var globalSessions *session.SManager



	http.HandleFunc("/", login)
	http.HandleFunc("/secretPage", secretPage)
	http.ListenAndServe(":8080", nil)
	// navigate to http://localhost:8080/
}

