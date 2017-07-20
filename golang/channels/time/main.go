package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("+100ms.")
		case <-boom:
			fmt.Println("=500ms!")
			return
		}
	}
}
