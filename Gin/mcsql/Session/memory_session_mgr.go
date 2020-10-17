package Session

import (
	"sync"
	uuid "github.com/satori/go.uuid"
)

type MemorySessionMgr struct {
	sessionMap map[string]Session
	rwlock sync.RWMutex
}

func NewMemoryMessionMgr() *MemorySessionMgr {
	sr := &MemorySessionMgr{
		sessionMap: make(map[string]Session,1024),
	}
	return sr
}

func (m *MemorySessionMgr) Init(addr string, options ... string) error{
	return nil
}
func (m *MemorySessionMgr) CreateSession() (Session, error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()

	uuid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	return NewMemorySession(uuid.String()) ,nil
}
func (m *MemorySessionMgr) Get(sessionId string) (Session, error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	session, ok := m.sessionMap[sessionId]
	if !ok {
		return nil,nil
	}
	return session,nil
}