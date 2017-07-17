package main

import (
	"fmt"
	"time"
)

// Item , item of each element of array
type Item int

// BubbleSort , bubble sort algorithm
func BubbleSort(list []Item) {
	for i := 0; i < len(list); i++ {
		innerSort(list)
	}
}

func innerSort(list []Item) {
	fIndex := 0
	lIndex := 1
	for lIndex < len(list) {
		fItem := list[fIndex]
		lItem := list[lIndex]

		if fItem > lItem {
			list[fIndex] = lItem
			list[lIndex] = fItem
		}

		fIndex++
		lIndex++
	}
}

// BubbleSortOptimized , bubble sort algorithm
func BubbleSortOptimized(list []Item) {
	for i := 0; i < len(list); i++ {
		swap := innerSortOptimized(list, i)
		if !swap {
			return
		}
	}
}

func innerSortOptimized(list []Item, level int) bool {
	fIndex := 0
	lIndex := 1
	swap := false

	for lIndex < len(list)-level {
		fItem := list[fIndex]
		lItem := list[lIndex]

		if fItem > lItem {
			list[fIndex] = lItem
			list[lIndex] = fItem
			swap = true
		}

		fIndex++
		lIndex++
	}
	return swap
}

func main() {
	copy := []Item{5, 3, 6, 1, 8, 7, 10, 203, 22, 12389, 12, 88, 33, 82, 91, 0}
	firstTime := time.Now().UnixNano()
	BubbleSort(copy)
	fmt.Printf("time  : %d nano second(s)\n", time.Now().UnixNano()-firstTime)

	copy = []Item{5, 3, 6, 1, 8, 7, 10, 203, 22, 12389, 12, 88, 33, 82, 91, 0}
	firstTime = time.Now().UnixNano()
	BubbleSortOptimized(copy)
	fmt.Printf("time  : %d nano second(s)\n", time.Now().UnixNano()-firstTime)
}
