package binaryTree

import "fmt"

// Node represents the components of a binary search tree
type Node struct {
	Data  *int
	Left  *Node
	Right *Node
}

// Insert method appends a new node to left or right of parent node. The Data to add should not be already in tree.
func (n *Node) Insert(val int) {
	if n.Data == nil {
		n.Data = &val
	}
	if val > *n.Data {
		//Move right
		if n.Right == nil {
			n.Right = &Node{Data: &val}
		} else {
			n.Right.Insert(val)
		}

	} else if val < *n.Data {
		//Move left
		if n.Left == nil {
			n.Left = &Node{Data: &val}
		} else {
			n.Left.Insert(val)
		}
	}
}

// Search method will take in a Data value and RETURN true if there is a node with that value
func (n *Node) Search(val int) bool {
	if n == nil {
		return false
	}

	if val > *n.Data {
		//Move right
		return n.Right.Search(val)
	} else if val < *n.Data {
		//Move left
		return n.Left.Search(val)
	}

	return true
}

func (n *Node) PrintInOrder() {
	if n == nil {
		return
	}
	if n.Left != nil {
		n.Left.PrintInOrder()
	}
	fmt.Println(*n.Data)
	if n.Right != nil {
		n.Right.PrintInOrder()
	}
}
