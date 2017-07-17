package main

import (
	"fmt"
)

// Stack , stack struct
type Stack struct {
	Count int
	Top   *Node
}

// Node , stack element
type Node struct {
	Value interface{}
	Next  *Node
}

// Pop , pop from stack
func (s *Stack) Pop() *Node {
	if s.IsEmpty() {
		return nil
	}

	top := s.Top
	s.Top = top.Next
	s.Count--
	return top
}

// Push , push to stack
func (s *Stack) Push(item interface{}) {
	node := &Node{Value: item}
	if s.Top != nil {
		node.Next = s.Top
	}
	s.Top = node
	s.Count++
}

// Size , size of stack
func (s *Stack) Size() int {
	return s.Count
}

// Print , print the stack
func (s *Stack) Print() {
	fmt.Print("{")

	top := s.Top
	for top != nil {
		fmt.Print(top.Value)
		if top.Next != nil {
			fmt.Print(",")
		}
		top = top.Next
	}

	fmt.Print("}\n")
}

// IsEmpty , return true if size is zero
func (s *Stack) IsEmpty() bool {
	return s.Count == 0
}

func main() {
	s := &Stack{}
	s.Push(10)
	s.Push(20)
	s.Print()
	s.Pop()
	s.Print()
}
