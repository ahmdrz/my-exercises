package main

import (
	"fmt"
	"time"
)

// LinearSearch , search over an array
func LinearSearch(list []interface{}, key interface{}) int {
	for i, item := range list {
		if item == key {
			return i
		}
	}
	return -1
}

func main() {
	firstTime := time.Now().UnixNano()

	list := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}
	index := LinearSearch(list, 4)
	fmt.Printf("index : %d\n", index)

	fmt.Printf("time : %d\n", time.Now().UnixNano()-firstTime)
}
