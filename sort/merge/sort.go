package main

import (
	"fmt"
)

// Item , type of each element
type Item int

// MergeSort , merge sort algorithm.
func MergeSort(list []Item) []Item {
	if len(list) < 2 {
		return list
	}
	middle := len(list) / 2
	left := MergeSort(list[:middle])
	right := MergeSort(list[middle:])

	return merge(left, right)
}

func merge(left, right []Item) []Item {
	slice := make([]Item, 0, len(right)+len(left))
	for len(right) > 0 || len(left) > 0 {
		if len(left) == 0 {
			return append(slice, right...)
		}
		if len(right) == 0 {
			return append(slice, left...)
		}

		if left[0] < right[0] {
			slice = append(slice, left[0])
			left = left[1:]
		} else {
			slice = append(slice, right[0])
			right = right[1:]
		}
	}
	return left
}

func main() {
	list := []Item{1, 2, 81, 91, 291, 129871, 1233, 501, 123, 81, 692, 23581, 811, 811}
	list = MergeSort(list)
	fmt.Println(list)
}
