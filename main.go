package main

import (
	"fmt"
	"redis/store/kv"
)

func main() {
	s := kv.NewStore(0)
	s.Set("pageviews", "0")
	const hits = 1000
	cmds := make([]Command, hits)
	for i := range cmds {
		cmds[i] = Command{Op: "INCR", Key: "pageviews"}
	}
	RestoreOnBoot(s, cmds)
	got, _ := s.Get("pageviews")

	fmt.Println("expected:", hits, "got:", got)
}
