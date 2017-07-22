package main

import (
	"fmt"

	_ "github.com/ahmdrz/my-exercises/golang/init/library/lib"
)

func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println("main")
}
