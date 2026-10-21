package ttl

import (
	"fmt"
	"time"

	"redis/store"
)

type TTLStore struct {
	data map[string]ttlEntry
	ttl  time.Duration
}

type ttlEntry struct {
	value        string
	expirationAt time.Time
}

func NewTTLStore(ttl time.Duration) *TTLStore {
	return &TTLStore{
		data: make(map[string]ttlEntry),
		ttl:  ttl,
	}
}

func (s *TTLStore) Set(key, value string) error {
	if key == "" {
		return store.ErrEmptyKey
	}

	expiration := time.Now().Add(s.ttl)
	s.data[key] = ttlEntry{
		value:        value,
		expirationAt: expiration,
	}
	return nil
}

func (t *TTLStore) Get(key string) (string, error) {
	if key == "" {
		return "", store.ErrEmptyKey
	}
	entry, ok := t.data[key]
	if !ok || time.Now().After(entry.expirationAt) {
		delete(t.data, key)
		return "", fmt.Errorf("key %s doesnt exist", key)
	}
	return entry.value, nil
}

func (t *TTLStore) Delete(key string) {
	delete(t.data, key)
}

func (t *TTLStore) Len() int {
	return len(t.data)
}
