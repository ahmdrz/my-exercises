package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var state = uint64(0)

	go func() {
		time.Sleep(1 * time.Second)
		atomic.AddUint64(&state, 1)
	}()

	for {
		current := atomic.LoadUint64(&state)
		fmt.Println("state : ", current)
		if current == 1 {
			return
		}

		time.Sleep(200 * time.Millisecond)
	}
}
