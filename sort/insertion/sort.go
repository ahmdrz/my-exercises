package main

import (
	"fmt"
	"time"
)

// Item , item of each element of array
type Item int

// InsertionSort , algorithm
func InsertionSort(list []Item) {
	for i := 1; i < len(list); i++ {
		j := i
		for j > 0 {
			current := list[j]
			lower := list[j-1]

			// check for sorted item , > for Ascending and < for Descending
			if current > lower {
				break
			}

			list[j] = lower
			list[j-1] = current
			j--
		}
	}
}

func main() {
	copy := []Item{5, 3, 6, 1, 8, 7, 10, 203, 22, 12389, 12, 88, 33, 82, 91, 0}
	firstTime := time.Now().UnixNano()
	InsertionSort(copy)
	fmt.Printf("time  : %d nano second(s)\n", time.Now().UnixNano()-firstTime)
	fmt.Println(copy)
}
