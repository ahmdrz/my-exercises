package main

import (
	"fmt"
)

func getEvenNumbers(a int) chan int {
	output := make(chan int)
	go func() {
		for i := 0; i < a; i += 2 {
			output <- i
		}
		close(output)
	}()
	return output
}

func main() {
	for x := range getEvenNumbers(10) {
		fmt.Println(x)
	}
}
