package main

import (
	"fmt"
	"sync"
)

var once sync.Once

var singleton map[string]int

func newMap() map[string]int {
	once.Do(func() {
		singleton = make(map[string]int)
	})

	singleton["count"]++
	return singleton
}

func main() {
	mapList1 := newMap()
	fmt.Println(mapList1["count"])
	mapList2 := newMap()
	fmt.Println(mapList2["count"])
}
