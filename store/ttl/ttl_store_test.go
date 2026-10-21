package ttl

import (
	"testing"
	"time"
)

func TestTTLStore_GetBeforeExpiring(t *testing.T) {
	s := NewTTLStore(1 * time.Second)
	_ = s.Set("key", "value")
	value, err := s.Get("key")
	if err != nil || value != "value" {
		t.Fatalf("expected v, got %q %v", value, err)
	}
}

func TestTTLStore_GetAfterExpiring(t *testing.T) {
	s := NewTTLStore(50 * time.Millisecond)
	_ = s.Set("key", "value")
	time.Sleep(100 * time.Millisecond)
	_, err := s.Get("key")
	if err == nil {
		t.Fatal("expected error after TTL, got nil")
	}
}

func TestTTLStore_ExpiredKeyEvicted(t *testing.T) {
	s := NewTTLStore(50 * time.Millisecond)
	_ = s.Set("key", "value")
	time.Sleep(100 * time.Millisecond)
	_, _ = s.Get("key")
	if s.Len() != 0 {
		t.Fatalf("expected 0 length after expired key eviction, got %d", s.Len())
	}
}
