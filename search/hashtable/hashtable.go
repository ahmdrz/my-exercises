package main

import "fmt"

const size = 1000

var hashtable []string = make([]string, size)

func hashCode(key string) int {
	value := 0
	for i := 0; i < len(key); i++ {
		value = 7*value + int(key[i])
	}
	return value
}

func generateIndex(key string) int {
	return hashCode(key) % size
}

func set(key, value string) {
	index := generateIndex(key)
	// ignore checking for duplicated index
	hashtable[index] = value
}

func get(key string) string {
	index := generateIndex(key)
	return hashtable[index]
}

func main() {
	set("ahmad", "reza")
	fmt.Println(get("ahmad"))
	fmt.Println(get("reza"))
}
