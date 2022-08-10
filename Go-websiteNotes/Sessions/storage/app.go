package main

import (

	"container/list"
	"github.com/astaxie/session"
	"sync"
	"time"
)
var globalSessions *session.SManager
var provides = make(map[string]Provider)
// Session manager

type SManager struct {
	cookieName string
	lock sync.Mutex
	provider Provider
	maxLifeTime int64
}

type Provider interface {
	SeshInit(SeshId string) (Session, error)
	SeshRead(SeshId string) (Session, error)
	SeshTerminate(SeshId string) error
	SeshLife(maxLifeTime int64)
}

type Session interface {
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Delete(key interface{}) error
	SeshionID() string
}

func NewManager(Name, cookieName string, maxLifeTime int64) (*SManager, error) {
	provider, ok := provides[Name]
	if !ok {
		return nil
	}
	return &SManager{provider: provider, cookieName: cookieName, maxLifeTime: maxLifeTime}, nil
}


func (manager *SManager) UseshID() string {
	// generate rand Uniqe id for new sessions
	UID := make([]byte,32)

	return base64.URLEncoding.EncodeToString(UID) 

}


func Register(name string, provider Provider) {
	if provider == nil {
		// panic
	}
	if _, dup := provides[name]; dup {
		// panic
	}
	provides[name] = provider
}

func (manager *SManager) SeshStart(w http.ResponseWriter, r *http.Request) (session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		SeshId := manager.UseshID()
		session, _ = manager.provider.SessionInit(SeshId)
		cookie := http.Cookie{Name: manager.cookieName, Value: url.QueryEscape(SeshId), Path: "/", HttpOnly: true, MaxAge: int(manager.maxLifeTime)}
		http.SetCookie(w, &cookie)
	} else {
		SeshId, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.provider.SessionRead(SeshId)
	}
	return
}

func (manager *SManager) GC() {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	manager.provider.SeshLife(manager.maxLifeTime)
	time.AfterFunc(time.Duration(manager.maxLifeTime), func() { manager.GC() })
}


func init() {
	globalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
	go globalSessions.GC()
}


// Protect from Session hijacking 
//solution 1 
//  only set session id's in cookies, instead of in URL rewrites
// httponly cookie is set to true.
// add a token to every request.
// add a hidden field that contains a token, so when a request is sent to the server, we can verify the token to prove request is unique.
//
// Solution2 
// create a time for every session, and replace expired session id's with new ones. 
// dose not protect from sessions that are currently active. 
