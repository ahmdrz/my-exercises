package main

import "fmt"

type item struct {
	id    int
	value int
}

type node struct {
	left  *node
	right *node
	i     item
}

type tree struct {
	root *node
	size int
}

func (t *tree) Size() int {
	return t.size
}

func (t *tree) Root() *node {
	return t.root
}

func newTree() *tree {
	bst := new(tree)
	bst.size = 0
	return bst
}

func (root *node) insert(new_node *node) {
	if new_node.i.id > root.i.id {
		if root.right == nil {
			root.right = new_node
		} else {
			root.right.insert(new_node)
		}
	} else if new_node.i.id < root.i.id {
		if root.left == nil {
			root.left = new_node
		} else {
			root.left.insert(new_node)
		}
	}
}

func (t *tree) Insert(id, value int) {
	if t.root == nil {
		t.root = &node{nil, nil, item{id, value}}
	}
	t.size++
	t.root.insert(&node{nil, nil, item{id, value}})
}

func showInOrder(root *node) {
	if root != nil {
		showInOrder(root.left)
		fmt.Printf("%d -> %d\n", root.i.id, root.i.value)
		showInOrder(root.right)
	}
}

func (t *tree) search(id int) *item {
	if t.root != nil {
		return t.root.Search(id)
	}
	return nil
}

func (n *node) Search(id int) *item {
	if n.i.id == id {
		return &n.i
	}
	if n.i.id < id {
		if n.right != nil {
			return n.right.Search(id)
		}
	}
	if n.i.id > id {
		if n.left != nil {
			return n.left.Search(id)
		}
	}
	return nil
}

func main() {
	t := newTree()
	t.Insert(1, 10)
	t.Insert(2, 20)
	t.Insert(3, 5)
	for j := 4; j < 40; j++ {
		t.Insert(j, j*3)
	}
	i := t.search(6)
	fmt.Println(t.Size())
	fmt.Println(i)
}
