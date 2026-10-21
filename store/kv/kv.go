package kv

import (
	"redis/store"
	"sort"
)

type Store struct {
	data    map[string]string
	maxKeys int
}

func NewStore(maxKeys int) *Store {
	return &Store{
		data:    make(map[string]string),
		maxKeys: maxKeys,
	}
}

// list keys in sorted order
func (s *Store) Keys() []string {
	keys := make([]string, 0, s.Len())
	for key := range s.data {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func (s *Store) Get(key string) (string, error) {
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
	_, exists := s.data[key]
	if s.Len() >= s.maxKeys && s.maxKeys > 0 && !exists {
		return store.ErrMaxKeysExceeded
	}
	s.data[key] = value
	return nil
}

func (s *Store) Delete(key string) {
	delete(s.data, key)
}

func (s *Store) Len() int {
	return len(s.data)
}
