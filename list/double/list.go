package main

import (
	"fmt"
)

// Node struct
type Node struct {
	Value interface{}
	Next  *Node
	Prev  *Node
}

// List struct
type List struct {
	Head *Node
	Last *Node
}

// Insert function
func (l *List) Insert(index int, item interface{}) {
	// I'm not going to check index with length...

	current := 0
	head := l.Head
	length := l.Length() - 1
	if index == 0 {
		l.InsertFirst(item)
		return
	}
	if index == length {
		l.InsertLast(item)
		return
	}

	node := &Node{Value: item}

	for head != nil {
		if current == index {
			node.Prev = head.Prev
			node.Next = head
			head.Next.Prev = node
			head.Prev.Next = node
			return
		}
		current++
		head = head.Next
	}
}

// InsertFirst function
func (l *List) InsertFirst(item interface{}) {
	node := &Node{Value: item}
	if l.IsEmpty() {
		l.Last = node
	} else {
		l.Head.Prev = node
	}
	node.Next = l.Head
	l.Head = node
}

// InsertLast function
func (l *List) InsertLast(item interface{}) {
	node := &Node{Value: item}
	if l.IsEmpty() {
		l.Last = node
	} else {
		l.Last.Next = node
		node.Prev = l.Last
	}
	l.Last = node
}

// Delete function
func (l *List) Delete(position int) {
	index := 0
	head := l.Head
	for head != nil {
		if index == position {
			if head == l.Head {
				l.Head = l.Head.Next
			} else {
				head.Prev.Next = head.Next
			}

			if head == l.Last {
				l.Last = l.Last.Prev
			} else {
				head.Next.Prev = head.Prev
			}
			return
		}
		index++
		head = l.Head.Next
	}
}

// Print function
func (l *List) Print() {
	head := l.Head
	fmt.Print("{")
	for head != nil {
		fmt.Print(head.Value)
		head = head.Next
		if head != nil {
			fmt.Print(",")
		}
	}
	fmt.Print("}\n")
}

// IsEmpty function
func (l *List) IsEmpty() bool {
	return l.Head == nil
}

// Length function
func (l *List) Length() int {
	length := int(0)
	head := l.Head
	for head != nil {
		length++
		head = head.Next
	}
	return length
}

func main() {
	list := List{}
	list.InsertFirst(10)
	list.InsertFirst(200)
	list.InsertLast(100)
	list.InsertLast(300)
	list.InsertLast(500)
	list.Delete(1) // item 10
	list.Insert(2, 15)
	list.Insert(2, 15)
	list.Print()
	fmt.Printf("Length : %d \n", list.Length())
}
