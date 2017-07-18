package main

import (
	"fmt"
	"sync"
	"time"
)

// Item , type of each element
type Item int

// MergeSort , merge sort algorithm.
func MergeSort(list []Item) []Item {
	if len(list) < 2 {
		return list
	}
	middle := len(list) / 2

	var left, right []Item

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		left = MergeSort(list[:middle])
		wg.Done()
	}()

	go func() {
		right = MergeSort(list[middle:])
		wg.Done()
	}()

	wg.Wait()
	return merge(left, right)
}

// SimpleMergeSort , old merge sort.
func SimpleMergeSort(list []Item) []Item {
	if len(list) < 2 {
		return list
	}
	middle := len(list) / 2

	var left, right []Item

	left = MergeSort(list[:middle])

	right = MergeSort(list[middle:])

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
	{
		list := []Item{1123123, 2, 81, 91, 291, 129871, 1233, 501, 123, 81, 692, 23581, 811, 811, 291, 129871, 1233, 501, 123, 81, 692, 23581, 811, 811}
		firstTime := time.Now().UnixNano()
		list = MergeSort(list)
		fmt.Printf("time  : %d nano second(s)\n", time.Now().UnixNano()-firstTime)
		fmt.Println(list)
	}

	{
		list := []Item{1123123, 2, 81, 91, 291, 129871, 1233, 501, 123, 81, 692, 23581, 811, 811, 291, 129871, 1233, 501, 123, 81, 692, 23581, 811, 811}
		firstTime := time.Now().UnixNano()
		list = SimpleMergeSort(list)
		fmt.Printf("time  : %d nano second(s)\n", time.Now().UnixNano()-firstTime)
		fmt.Println(list)
	}
}
