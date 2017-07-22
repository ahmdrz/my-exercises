package main

import "regexp"
import "fmt"

var re = regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}`)

func main() {
	matches := re.FindAllString("hello my email is ahmadreza@zibaei.net or (ahmadrezazibaei@hotmail.com)", -1)
	for _, email := range matches {
		fmt.Println(email)
	}
}
