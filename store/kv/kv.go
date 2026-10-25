package kv

import (
	"redis/store"
	"sort"
	"strconv"
	"sync"
)

type Store struct {
	mu      sync.RWMutex
	data    map[string]string
	maxKeys int
}

func NewStore(maxKeys int) *Store {
	return &Store{
		data:    make(map[string]string),
		maxKeys: maxKeys,
	}
}

func (s *Store) Clone() *Store {
	s.mu.RLock()
	defer s.mu.RUnlock()

	clone := &Store{
		data:    make(map[string]string, len(s.data)),
		maxKeys: s.maxKeys,
	}
	for key, value := range s.data {
		clone.data[key] = value
	}
	return clone
}

// list keys in sorted order
func (s *Store) Keys() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	keys := make([]string, 0, len(s.data))
	for key := range s.data {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func (s *Store) Get(key string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, ok := s.data[key]
	if !ok {
		return "", store.ErrEmptyKey
	}
	return val, nil
}

func (s *Store) Set(key, value string) error {
	if key == "" {
		return store.ErrEmptyKey
	}
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.data[key]
	if len(s.data) >= s.maxKeys && s.maxKeys > 0 && !exists {
		return store.ErrMaxKeysExceeded
	}
	s.data[key] = value
	return nil
}

func (s *Store) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.data, key)
}

func (s *Store) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return len(s.data)
}

func (s *Store) Incr(key string) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if key == "" {
		return 0, store.ErrEmptyKey
	}
	// Read current value
	val, ok := s.data[key]
	if !ok {
		return 0, store.ErrEmptyKey
	}
	// Increment
	num, err := strconv.Atoi(val)
	if err != nil {
		return 0, err
	}
	num++
	// Write new value
	s.data[key] = strconv.Itoa(num)
	return num, nil
}
