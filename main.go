package main

import "fmt"

func main() {
	s := NewStore(2)
	s.Set("a", "1")
	s.Set("b", "2")
	err := s.Set("3", "3")
	if err != nil {
		fmt.Printf("%v", err)
	}
	// val, err := s.Get("a")
	// delete(s.data, "a")
	// val2, err2 := s.Get("a")
	// fmt.Println(val, err)
	// fmt.Println(val2, err2)
}
