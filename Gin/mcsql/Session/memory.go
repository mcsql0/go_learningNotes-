package Session

import (
	"errors"
	"sync"
)

type MemorySession struct {
	sessionId string
	data map[string]interface{}
	rwlock sync.RWMutex
}

func NewMemorySession(id string) *MemorySession {
	s := &MemorySession{
		sessionId: id,
		data: make(map[string]interface{},16),
	}
	return s
}

func (m *MemorySession) Set(key string, value interface{}) error{
	m.rwlock.Lock()
	defer m.rwlock.Unlock()

	m.data[key] = value
	return nil
}
func (m *MemorySession) Get(key string) (interface{}, error){
	m.rwlock.Lock()
	defer m.rwlock.Unlock()

	v,ok := m.data[key]
	if !ok {
		return nil,errors.New("key 不存在")
	}
	return v,nil
}
func (m *MemorySession) Del(key string) error{
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	delete(m.data,key)
	return nil
}
func (m *MemorySession) Save() error{
	return nil
}
