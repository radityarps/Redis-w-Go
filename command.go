package main

import (
	"context"
	"errors"
	"fmt"
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
func dispatch(s store.Storer, c Command) error {
	switch c.Op {
	case "SET":
		return s.Set(c.Key, c.Value)
	case "GET":
		_, err := s.Get(c.Key)
		return err
	case "INCR":
		_, err := s.Incr(c.Key)
		return err
	default:
		return fmt.Errorf("unknown op: %s", c.Op)
	}
}

// RestoreOnBoot replays a slice of command against a store
func RestoreOnBoot(ctx context.Context, s store.Storer, cmds []Command) error {
	var wg sync.WaitGroup
	errs := make(chan error, len(cmds))

	for _, c := range cmds {
		wg.Go(
			func() {
				select {
				case <-ctx.Done():
					errs <- ctx.Err()
				default:
					errs <- dispatch(s, c)
				}
			},
		)
	}
	go func() {
		wg.Wait()
		close(errs)
	}()
	var allErrs []error

	for err := range errs {
		if err != nil {
			allErrs = append(allErrs, err)
		}
	}
	if len(allErrs) > 0 {
		return errors.Join(allErrs...)
	}
	return nil
}
