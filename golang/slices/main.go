package main

import (
	"fmt"
)

// T is type of int
type T int

func main() {
	{
		// copy
		listA := []T{1, 2, 3}
		listB := append([]T(nil), listA...)
		// or copy(listA, listB)

		fmt.Println("copy", listA, listB)
	}

	{
		// cut
		listA := []T{1, 2, 3, 4, 5}
		listA = append(listA[:2], listA[3:]...)

		fmt.Println("cut", listA)
	}

	{
		// delete
		listA := []T{1, 2, 3, 4, 5}
		listA = append(listA[:1], listA[1+1:]...)

		fmt.Println("delete", listA)
	}

	{
		// extend
		listA := []T{1, 2, 3}
		listB := []T{4, 5, 6}
		listA = append(listA, listB...)

		fmt.Println("extend", listA)
	}

	{
		// insert
		listA := []T{1, 2, 3}
		listA = append(listA, 0)

		// insert at 1
		i := 1

		copy(listA[i+1:], listA[i:])
		listA[i] = 10

		fmt.Println("insert", listA)
	}

	{
		// change on for
		listA := []T{1, 2, 3}
		for _, item := range listA {
			listA = append(listA, item)
		}

		fmt.Println("change on for", listA)
	}
}
