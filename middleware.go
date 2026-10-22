package main

import (
	"log"
	"os"
	"redis/store"
)

type LoggingMiddleware struct {
	inner  store.Storer
	logger *log.Logger
}

func NewLoggingMiddleware(inner store.Storer) *LoggingMiddleware {
	return &LoggingMiddleware{
		inner:  inner,
		logger: log.New(os.Stdout, "[log]", log.Ltime|log.Lmicroseconds),
	}
}

func (m *LoggingMiddleware) Get(key string) (string, error) {
	m.logger.Printf("GET %q\n", key)
	value, err := m.inner.Get(key)
	if err != nil {
		m.logger.Printf("GET %q ERROR: %v\n", key, err)
	}
	return value, err
}

func (m *LoggingMiddleware) Set(key string, value string) error {
	m.logger.Printf("SET %q %s\n", key, value)
	err := m.inner.Set(key, value)
	return err
}

func (m *LoggingMiddleware) Delete(key string) {
	m.logger.Printf("DELETE %q\n", key)
	m.inner.Delete(key)
}

func (m *LoggingMiddleware) Len() int {
	len := m.inner.Len()
	m.logger.Printf("LEN %d\n", len)
	return len
}

func (m *LoggingMiddleware) Keys() []string {
	keys := m.inner.Keys()
	m.logger.Printf("KEYS %v\n", keys)
	return keys
}
