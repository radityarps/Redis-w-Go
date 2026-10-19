package main

import (
	"fmt"
)

func main() {
	s := NewStore()
	s.Set("a", "1")
	val, err := s.Get("a")
	delete(s.data, "a")
	val2, err2 := s.Get("a")
	fmt.Println(val, err)
	fmt.Println(val2, err2)
}
