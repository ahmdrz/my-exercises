package main

import (
	"fmt"
)

func findGoodResult(money int, ourCoins []int) int {
	for _, item := range ourCoins {
		if item <= money {
			return item
		}
	}
	return 1
}

func calculateCoins(money int, ourCoins []int) []int {
	output := make([]int, 0)
	for money != 0 {
		item := findGoodResult(money, ourCoins)
		output = append(output, item)
		money -= item
	}
	return output
}

func main() {

	// best implement is using with maps
	// btw , I used arrays and slices

	money := 72
	ourCoins := []int{25, 15, 10, 5, 1}

	for _, item := range calculateCoins(money, ourCoins) {
		fmt.Println(item)
	}
}
