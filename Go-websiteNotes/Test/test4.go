package main

import (
	"net/http"
	"github.com/gorilla/sessions"
)


var store = sessions.NewCookieStore(
	[]byte("new-Auth-key"),
	[]byte("new-encryption-key"),
	[]byte("old-auth-key"),
	[]byte("old-encryption-key")
)


// Session manager


type SManager struct {
	cookieName	string // private 
	lcok		sync.Mutex // protects session
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
	SessionID() string	
}

func login(w http.ResponseWriter, r *http.Request) {
	sesh := globalSessions.SessionStart(w, r)
	

	r.ParseForm()
	if r.Method == "GET"{
		tmpl, _ :=template.ParseFiles("login.gtpl")
		tmpl.Execute(w, sesh.Get("username"))
	}

}

func NewManager(givename, cookieName string, maxlifetime int64) (*SManager, error) {
	provider, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("Session: unkown provide %q (forgotten import?)", givename)
	}
	return &Manager{provider: provider, cookieName: cookieName, maxlifetime: maxlifetime}, nil
}

func init() {
	globalSessions = NewManager("memory", "gosessionid", 3600)
}


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
		session, _ = manaer.provider.SessionRead(sid)
	}
	return
}

func main() {


	// create global session manager 
	var globalSessions *session.Manager

}
