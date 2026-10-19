package main

import (
	"fmt"
)


func main() {
	s := NewStore()
	s.Set("a", "1")
	val, ok := s.Get("a")
	delete(s.data, "a")
	val2, ok2 := s.Get("a")
	fmt.Println(val, ok)
	fmt.Println(val2, ok2)
}
