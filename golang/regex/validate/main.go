package main

import (
	"fmt"
	"regexp"
)

func validateEmail(email string) bool {
	return regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`).MatchString(email)
}

func main() {
	email := "ahmadreza@zibaei.net"
	fmt.Println(email, validateEmail(email))
}
