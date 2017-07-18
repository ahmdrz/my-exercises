package main

import (
	"fmt"
)

type deviderType int

const (
	devideOne deviderType = 1 << iota
	devideTwo
	devideThree
)

type devider interface {
	Devide(a int) int
}

type deviderStruct struct {
	devider
	devide int
}

func (d deviderStruct) Devide(a int) int {
	return a / d.devide
}

func newDevider(t deviderType) devider {
	switch t {
	case devideOne:
		return deviderStruct{devide: 1}
	case devideTwo:
		return deviderStruct{devide: 2}
	default:
		return deviderStruct{devide: 3}
	}
}

func main() {
	d := newDevider(devideThree)
	fmt.Println(d.Devide(99))
}
