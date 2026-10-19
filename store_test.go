package main

import (
	"errors"
	"slices"
	"testing"
)

func TestStore_Keys(t *testing.T) {
	t.Run("returns empty slice when store is empty", func(t *testing.T) {
		s := NewStore()
		keys := s.Keys()

		if len(keys) != 0 {
			t.Errorf("expected 0 keys, got %d", len(keys))
		}
	})

	t.Run("returns single key", func(t *testing.T) {
		s := NewStore()
		s.Set("key1", "val1")

		keys := s.Keys()
		expected := []string{"key1"}

		if !slices.Equal(keys, expected) {
			t.Errorf("expected %v, got %v", expected, keys)
		}
	})

	t.Run("returns keys in ascending sorted order", func(t *testing.T) {
		s := NewStore()
		s.Set("banana", "1")
		s.Set("apple", "2")
		s.Set("date", "3")
		s.Set("cherry", "4")

		keys := s.Keys()
		expected := []string{"apple", "banana", "cherry", "date"}

		if !slices.Equal(keys, expected) {
			t.Errorf("expected %v, got %v", expected, keys)
		}
	})

	t.Run("returns sorted keys correctly after deletion", func(t *testing.T) {
		s := NewStore()
		s.Set("b", "1")
		s.Set("a", "2")
		s.Set("c", "3")
		s.Delete("b")

		keys := s.Keys()
		expected := []string{"a", "c"}

		if !slices.Equal(keys, expected) {
			t.Errorf("expected %v, got %v", expected, keys)
		}
	})
}

func TestSetGet_RoundTrip(t *testing.T) {
	s := NewStore()
	s.Set("key1", "val1")

	if val, err := s.Get("key1"); err != nil || val != "val1" {
		t.Errorf("expected key to be found with value 'val1', got '%s' %v", val, err)
	}

	if val, err := s.Get("missing"); err == nil || val != "" {
		t.Errorf("expected key to be not found, got '%s' %v", val, err)
	}
}

func TestSetGet_EmptyKeys(t *testing.T) {
	s := NewStore()
	if _, err := s.Get(""); err == nil && !errors.Is(err, ErrEmptyKey) {
		t.Errorf("Get() failed: %v", err)
	}
}

// TableTests
func TestDelete(t *testing.T) {
	tests := []struct {
		name        string
		setup       map[string]string
		expectedLen int
		deleteKey   string
	}{
		{
			name: "delete existing key",
			setup: map[string]string{
				"key1": "val1",
				"key2": "val2",
			},
			expectedLen: 1,
			deleteKey:   "key1",
		},
		{
			name: "delete non-existing key",
			setup: map[string]string{
				"key1": "val1",
			},
			expectedLen: 1,
			deleteKey:   "x",
		},
		{
			name:        "delete on empty store",
			setup:       map[string]string{},
			expectedLen: 0,
			deleteKey:   "x",
		},
		{
			name: "delete last remaining key",
			setup: map[string]string{
				"key1": "val1",
			},
			expectedLen: 0,
			deleteKey:   "key1",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := NewStore()
			for k, v := range test.setup {
				s.Set(k, v)
			}
			s.Delete(test.deleteKey)
			if s.Len() != test.expectedLen {
				t.Errorf("expected %d keys, got %d", test.expectedLen, s.Len())
			}
			if _, err := s.Get(test.deleteKey); err == nil {
				t.Errorf("expected key to be deleted, got %v", err)
			}
		})
	}
}
