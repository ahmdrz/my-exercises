package main

import (
	"fmt"
)

// Item , element item
type Item int

// SelectionSort algorithm
func SelectionSort(list []Item) {
	for i := 0; i < len(list)-1; i++ {
		index := i
		for j := i + 1; j < len(list); j++ {
			if list[index] > list[j] {
				index = j
			}
		}
		if index != i {
			list[i], list[index] = list[index], list[i]
		}
	}
}

func main() {
	list := []Item{811, 1, 123, 165, 92, 192, 74, 58, 29, 71, 393, 293}
	SelectionSort(list)
	fmt.Println(list)
}
