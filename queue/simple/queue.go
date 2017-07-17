package main

import (
	"fmt"
)

// Queue , Queue struct
type Queue struct {
	Count int
	Front *Node
	Rear  *Node
}

// Node , Queue element
type Node struct {
	Value interface{}
	Next  *Node
}

// DeQueue , dequeue from Queue
func (q *Queue) DeQueue() *Node {
	if q.IsEmpty() {
		return nil
	}
	last := q.Front
	if q.Front.Next == nil {
		q.Front = nil
		q.Rear = nil
	} else {
		q.Front = q.Front.Next
	}
	return last
}

// EnQueue , enqueue to Queue
func (q *Queue) EnQueue(item interface{}) {
	node := &Node{Value: item}
	if q.IsEmpty() {
		q.Front = node
		q.Rear = node
	} else {
		q.Rear.Next = node
		q.Rear = node
	}
	q.Count++
}

// Size , size of Queue
func (q *Queue) Size() int {
	return q.Count
}

// Print , print the Queue
func (q *Queue) Print() {
	fmt.Print("{")

	top := q.Front
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
func (q *Queue) IsEmpty() bool {
	return q.Count == 0
}

func main() {
	q := &Queue{}
	q.EnQueue(10)
	q.EnQueue(20)
	q.Print()
	q.DeQueue()
	q.Print()
}
