package main

import (
	"context"
	"fmt"
	"redis/store/kv"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	defer cancel()
	s := kv.NewStore(0)
	s.Set("pageviews", "0")
	const hits = 1000
	cmds := make([]Command, hits)
	for i := range cmds {
		cmds[i] = Command{Op: "INCR", Key: "pageviews"}
	}

	cmds = append(cmds, Command{Op: "INCR", Key: ""}, Command{Op: "WAT", Key: "pageviews"})

	if err := RestoreOnBoot(ctx, s, cmds); err != nil {
		fmt.Printf("Restore Aborted: %v\n", err)
	}
	got, _ := s.Get("pageviews")

	fmt.Println("expected:", hits, "got:", got)
}
