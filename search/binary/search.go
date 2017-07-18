package main

import (
	"fmt"
	"time"
)

// BinarySearch , search over an array
func BinarySearch(list []int, key int) int {
	low, high := 0, len(list)-1
	for low <= high {
		mid := (low + high) / 2

		if list[mid] < key {
			low = mid + 1
		} else if list[mid] > key {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	firstTime := time.Now().UnixNano()

	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	index := BinarySearch(list, 9)
	fmt.Printf("index : %d\n", index)

	fmt.Printf("time : %d\n", time.Now().UnixNano()-firstTime)
}
