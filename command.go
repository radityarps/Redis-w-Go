package main

import (
	"redis/store"
	"sync"
)

// Command represents a parsed Redis-ish operation: opcode + its arguments
type Command struct {
	Op    string
	Key   string
	Value string
}

// dispatch is a tiny switch that turns a Command into a store cell
func dispatch(s store.Storer, c Command) {
	switch c.Op {
	case "SET":
		s.Set(c.Key, c.Value)
	case "GET":
		s.Get(c.Key)
	}
}

// RestoreOnBoot replays a slice of command against a store
func RestoreOnBoot(s store.Storer, cmds []Command) {
	var wg sync.WaitGroup
	for _, c := range cmds {
		wg.Go(
			func() {
				dispatch(s, c)
				wg.Done()
			},
		)
	}
	wg.Wait()
}
