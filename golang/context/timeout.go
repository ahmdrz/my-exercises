package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("+100ms.")
		case <-boom:
			fmt.Println("=500ms!")
			return
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		}
	}
}
