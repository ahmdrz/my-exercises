package main

import (
	"fmt"
)

// don't care about is array thread-safe or not...

var memory = []int{0, 1}

func fib(a int) int {
	if a < len(memory) {
		return memory[a]
	}
	tmp := fib(a-1) + fib(a-2)
	memory = append(memory, tmp)
	return tmp
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(fib(i))
	}
}
