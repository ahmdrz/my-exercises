package main

import "fmt"
import "time"
import "strings"

// RULES :
// Unlike Adapter pattern, the object to be decorated is obtained by injection.
// Decorators should not alter the interface of an object.

func timeCalculator(fn func(string) string) func(string) string {
	return func(n string) string {
		firstTime := time.Now().UnixNano()
		result := fn(n)
		fmt.Printf("time for %s is %d\n", n, time.Now().UnixNano()-firstTime)
		return result
	}
}

func parseString(s string) string {
	list := strings.Split(s, " ")
	return strings.Join(list, "+")
}

func main() {
	ts := timeCalculator(parseString)
	ts("hello world")
	ts("My name is Ahmadreza")
	ts("a b c d e f g h")
}
