package memory

import (

	"container/list"
	"github.com/astaxie/session"

	"sync"
	"time"
	_ "github.com/astaxie/session/providers/memory"
)




// Session storage

var PD = &Provider{list: list.New()}


type SessionStore struct {
	SeshId	string
	timeAccessed time.Time 
	value map[interface{}]interface{}
}

type Provider struct {
	lock sync.Mutex
	sessions map[string]*list.Element
	list *list.List
}


func (storage *SessionStore) Set(key, value interface{}) error {
	storage.value[key] = value
	// test this
	PD.SeshUpdate(storage.SeshId)

	return nil
}

func (storage *SessionStore) Get(key interface{}) interface{} {
	PD.SeshUpdate(storage.SeshId)
	if V, ok := SeshId.Value[key]; ok {
		return V
	} else {
		return nil
	}
	return nil
}

func (storage *SessionStore) Delete(key interface{}) error {
	delete(storage.value, key)
	PD.SeshUpdate(storage.SeshId)
	return nil
}

func (storage *SessionStore) SeshionID() string {
	return storage.SeshId
}


func (PD *Provider) SeshInit(SeshId string) (session.Session, error) {
	PD.lock.Lock()
	defer PD.lock.Unlock()
	Val := make(mmap[interface{}]interface{}, 0)
	NewSesh := &SessionStore{SeshId: SeshId, timeAccessed: time.Now(), value: Val}
	element := PD.list.PushBack(NewSesh)
	PD.sessions[SeshId] = element
	return NewSesh, nil
}

func (PD *Provider) SeshRead(SeshId string) (session.Session, error) {
	if element, ok := PD.sessions[SeshId]; ok {
		return element.Value.(*SessionStore), nil
	}else {
		Sesh, err := PD.SeshInit(SeshId)

		return sesh, err

	}
	return nil, nil

}

func (PD *Provider) SeshTerminate(SeshId string) error{
	if element, ok := PD.sessions[SeshId]; ok {
		delete(PD.sessions, SeshId)
		PD.list.Remove(element)
		return nil
	}
	return nil

}

func (PD *Provider) SeshLife(maxLifeTime int64) {
	PD.lock.Lock()
	defer.lock.Unlock()

	for {
		element := PD.list.Back()
		if element == nil {
			break
		}
		if (element.Value.(*SessionStore).timeAccessed.Unix() + maxLifeTime) < time.Now().Unix() {
			PD.list.Remove(element)
			delete(PD.sessions, element.Value.(*SessionStore).SeshId)
		} else {
			break
		}
	}
}

func (PD *Provider) SeshUpdate(SeshId string) error {
	PD.lock.Lock()
	defer PD.lock.Unlock()

	if element, ok := PD.sessions[SeshId]; ok {
		element.Value.(*SessionStore).timeAccessed = time.Now()
		PD.list.MoveToFront(element)
		return nil
	}
	return nil
}

func init() {
	PD.sessions = make(map[string]*list.Element, 0)
	session.Register("memory", PD)
}


