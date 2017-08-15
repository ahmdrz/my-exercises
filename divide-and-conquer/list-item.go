package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	// DEBUG , if true , program will process steps
	DEBUG = false
)

func detectWithoutWg(list []int, min, max int) int {
	if DEBUG {
		fmt.Println(min, max, list, max-min)
	}
	result := -1
	for i := 0; i < max-min; i++ {
		if DEBUG {
			fmt.Println(i, min+i, list[i])
		}
		if list[i] != min+i {
			result = list[i] - 1
			break
		}
	}
	if DEBUG {
		fmt.Println("output ----", result)
	}
	return result
}

func detectUsingWg(wg *sync.WaitGroup, list []int, min, max int, output chan int) {
	result := detectWithoutWg(list, min, max)
	if result != -1 {
		output <- result
	}
	wg.Done()
}

func detectUnusualNumberFastWithGo(number []int, parts int) int {
	wg := &sync.WaitGroup{}
	output := make(chan int)

	l := len(number)
	wg.Add(parts)
	l = l / parts
	for i := 0; i < parts; i++ {
		if i == 0 {
			go detectUsingWg(wg, number[0:l], 0, l, output)
		} else if i == parts-1 {
			go detectUsingWg(wg, number[l*i:], l*i, len(number), output)
		} else {
			min := l * i
			max := l * (i + 1)
			go detectUsingWg(wg, number[min:max], min, max, output)
		}
	}

	go func() {
		wg.Wait()
		close(output)
	}()
	result := <-output
	return result
}

func detectUnusualNumberFastWithoutGo(number []int, parts int) int {
	l := len(number)
	l = l / parts

	output := -1

	for i := 0; i < parts; i++ {
		if i == 0 {
			result := detectWithoutWg(number[0:l], 0, l)
			if result != -1 {
				output = result
				break
			}
		} else if i == parts-1 {
			result := detectWithoutWg(number[l*i:], l*i, len(number))
			if result != -1 {
				output = result
				break
			}
		} else {
			min := l * i
			max := l * (i + 1)
			result := detectWithoutWg(number[min:max], min, max)
			if result != -1 {
				output = result
				break
			}
		}
	}

	return output
}

func main() {
	rand.Seed(time.Now().UnixNano())

	list := make([]int, 0)

	a := 100
	b := 999
	length := a + rand.Intn(b-a) + 1
	time.Sleep(10 * time.Millisecond)
	random := rand.Intn(length) + 1

	for i := 0; i < length; i++ {
		if i == random {
			continue
		}
		list = append(list, i)
	}

	t := time.Now()
	result := detectUnusualNumberFastWithGo(list, 3)
	if DEBUG {
		fmt.Println(time.Since(t), result, list)
	} else {
		fmt.Println(time.Since(t), result)
	}
}
