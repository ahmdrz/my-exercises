package main

import (
	"fmt"
)

type memento struct {
	money int
}

func (self *memento) getMoney() int {
	return self.money
}

type game struct {
	Money int
}

func (self *game) createMemento() *memento {
	return &memento{
		self.Money,
	}
}

func (self *game) restoreMemento(memento *memento) {
	self.Money = memento.getMoney()
}

func main() {
	g := game{10}
	snapshot := g.createMemento()
	g.restoreMemento(snapshot)
	fmt.Println(g.Money)
}
