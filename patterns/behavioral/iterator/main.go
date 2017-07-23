package main

import (
	"fmt"
)

type iterator struct {
	position int
	data     []int
}

func (i *iterator) next() bool {
	if i.position >= len(i.data) {
		return false
	}
	return true
}

func (i *iterator) value() int {
	output := i.data[i.position]
	i.position++
	return output
}

func main() {
	list := []int{1, 2, 3, 4, 5}
	iter := &iterator{data: list}

	for iter.next() {
		fmt.Println(iter.value())
	}
}
