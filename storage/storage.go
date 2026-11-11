package storage

import "sync"

type Storage struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) Get(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.data[key]
}

func (s *Storage) Set(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}
