package main

import "fmt"

// Node represents the components of a binary search tree
type Node struct {
	Key int
	Left *Node
	Right *Node
}

func main() {
	tree := Node{Key: 89}
	tree.Insert(22)
	tree.Insert(190)
	tree.Insert(54)
	tree.Insert(647)
	tree.Insert(831)
	tree.Insert(656)

	fmt.Println(tree)
	fmt.Println(tree.Right.Left)
	fmt.Println(tree.Left)
}

// Insert method appends to new node to left or right of parent node. The key to add should not be already in tree.
func (n *Node) Insert(val int) {
	if val > n.Key {
		//Move right
		if n.Right == nil {
			n.Right = &Node{Key: val}
		} else {
			n.Right.Insert(val)
		}

	} else if val < n.Key {
		//Move left
		if n.Left == nil {
			n.Left = &Node{Key: val}
		} else {
			n.Left.Insert(val)
		}
	}
}


// Search method will take in a key value and RETURN true if there is a node with that value
func (n *Node) Search(val int) bool {
	if n == nil {
		return false
	}

	if val > n.Key {
		//Move right
		return n.Right.Search(val)
	} else if val < n.Key {
		//Move left
		return n.Left.Search(val)
	}

	return true
}