package main

import (
	"fmt"
	"redis/store/kv"
)

func main() {
	cmds := []Command{
		{Op: "SET", Key: "env", Value: "production"},
		{Op: "SET", Key: "ver", Value: "1.0.0"},
		{Op: "SET", Key: "debug", Value: "true"},
		{Op: "GET", Key: "env"},
		{Op: "SET", Key: "region", Value: "eu-west-1"},
		{Op: "GET", Key: "debug"},
	}

	s := kv.NewStore(0)
	RestoreOnBoot(s, cmds)
	fmt.Println("restored, store size:", s.Len())
	fmt.Println("keys:", s.Keys())

}
