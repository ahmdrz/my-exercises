package main

import (
	"fmt"
)

// Node struct
type Node struct {
	Value interface{}
	Next  *Node
}

// List struct
type List struct {
	Head *Node
}

// Insert function
func (l *List) Insert(item interface{}) {
	node := &Node{Value: item}
	if l.Head == nil {
		l.Head = node
		return
	}
	node.Next = l.Head
	l.Head = node
}

// Remove function
func (l *List) Remove(item interface{}) {
	if l.Head.Value == item {
		if l.Head.Next == nil {
			l.Head = nil
			return
		}
		l.Head = l.Head.Next
		return
	}
	head := l.Head
	for head != nil {
		if head.Next != nil {
			if head.Next.Value == item {
				if head.Next.Next == nil {
					head.Next = nil
					return
				}
				head.Next = head.Next.Next
				return
			}
		}
		head = head.Next
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
func (l *List) Length() uint {
	length := uint(0)
	head := l.Head
	for head != nil {
		length++
		head = head.Next
	}
	return length
}

func main() {
	list := List{}
	list.Insert(10)
	list.Insert(100)
	list.Remove(10)
	list.Insert(10)
	list.Remove(100)
	list.Remove(10)
	list.Print()
	fmt.Printf("Length : %d \n", list.Length())
}
