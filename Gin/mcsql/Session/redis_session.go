package Session

import (
	"encoding/json"
	"errors"
	"github.com/garyburd/redigo/redis"
	"sync"
)

const (
	SessionFlagNone = iota
	SessionFlagModify
)

type RedisSession struct {
	sessionid string
	pool *redis.Pool
	sessionMap map[string]interface{}
	rwlock sync.Locker
	flag int
}

func NewRedisSession(id string, pool *redis.Pool) *RedisSession  {
	s := &RedisSession{
		sessionid: id,
		sessionMap: make(map[string]interface{},16),
		pool: pool,
		flag: SessionFlagNone,
	}
	return s
}

func (r *RedisSession) Set(key string, value interface{}) error {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()

	r.sessionMap[key] = value
	r.flag = SessionFlagModify
	return nil
}

func (r *RedisSession) Get(key string) (interface{}, error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()

	result,ok := r.sessionMap[key]
	if !ok {
		return nil, errors.New("key not exists")
	}

	return result, nil
}

func (r *RedisSession) loadFromReids() error {
	conn := r.pool.Get()
	reply, err := conn.Do("GET", r.sessionid)
	if err != nil {
		return err
	}
	data, err := redis.String(reply, err)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(data), &r.sessionMap)
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisSession) Del(key string) error {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()

	r.flag = SessionFlagModify
	delete(r.sessionMap,key)
	return nil
}

func (r *RedisSession) Save() error {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()

	if r.flag != SessionFlagModify {
		return nil
	}
	data, err := json.Marshal(r.sessionMap)
	if err !=  nil {
		return err
	}

	//获取redis连接
	conn := r.pool.Get()
	_, err = conn.Do("SET", r.sessionid, string(data))
	if err !=  nil {
		return err
	}
	r.flag = SessionFlagNone
	return nil
}