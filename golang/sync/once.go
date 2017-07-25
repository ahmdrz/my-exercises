package main

import (
	"fmt"
	"sync"
)

var once = sync.Once{}

func sayHelloOnce() {
	once.Do(func() {
		fmt.Println("Hello World Once")
	})
}

func sayHello() {
	fmt.Println("Hello World")
}

func main() {
	// bayad yebar hello world chap kone
	// dar hali ke man 10 bar sedash mizanam
	for i := 0; i < 10; i++ {
		sayHelloOnce()
	}

	// hala bayad 10 bar hello world chap kone
	for i := 0; i < 10; i++ {
		sayHello()
	}
}
