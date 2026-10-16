package main

import (
	"fmt"
)

type Store struct {
	data map[string]string
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

func (s *Store) Get(key string) (string, bool) {
	val, ok := s.data[key]
	return val, ok
}

func (s *Store) Set(key, value string) {
	s.data[key] = value
}

func (s *Store) Delete(key string) {
	delete(s.data, key)
}

func main() {
	s := NewStore()
	s.Set("a", "1")
	val, ok := s.Get("a")
	delete(s.data, "a")
	val2, ok2 := s.Get("a")
	fmt.Println(val, ok)
	fmt.Println(val2, ok2)
}
