package redis

import (
	"sync"
)

type Redis struct {
	storage map[string]string
	mu      sync.RWMutex
}

func NewStorage() *Redis {
	return &Redis{
		storage: make(map[string]string),
	}
}

func (s *Redis) Set(key string, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.storage[key] = value

}

func (s *Redis) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	v, ok := s.storage[key]

	return v, ok
}

func (s *Redis) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.storage, key)
}
