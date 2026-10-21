package store

import "errors"

var ErrEmptyKey = errors.New("key cannot be empty")
var ErrMaxKeysExceeded = errors.New("max keys exceeded")

type Storer interface {
	Get(key string) (string, error)
	Set(key, value string) error
	Delete(key string)
	Len() int
	// Keys() []string
}
