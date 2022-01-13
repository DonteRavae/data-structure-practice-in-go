package binaryTree

import "fmt"

// Node represents the components of a binary search tree
type Node struct {
	Data  *int
	Left  *Node
	Right *Node
}

// Insert appends a new node to left or right of parent node. The Data to add should not be already in tree.
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

// Search will take in a Data value and RETURN true if there is a node with that value
func (n Node) Search(val int) bool {
	if n.Data == nil {
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

// PrintInOrder takes a depth-first traversal approach and prints the data in order if the tree is a binary search tree.
func (n Node) PrintInOrder() {
	if n.Data == nil {
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

// PrintLevelOrder takes a breadth-first traversal approach and prints the data starting from the root, moving left to right on each level of the tree.
func (n Node) PrintLevelOrder() {
	if n.Data == nil {
		return
	}
	queue := []Node{n}
	for len(queue) > 0 {
		node := queue[0]
		if node.Left != nil {
			queue = append(queue, *node.Left)
		}
		if node.Right != nil {
			queue = append(queue, *node.Right)
		}
		queue = queue[1:]
		fmt.Println(*node.Data)
	}
}

// Size returns the amount of nodes in the tree.
func (n Node) Size() int {
	if n.Data == nil {
		return 0
	}
	var sum int
	queue := []Node{n}
	for len(queue) > 0 {
		node := queue[0]
		sum++
		if node.Left != nil {
			queue = append(queue, *node.Left)
		}
		if node.Right != nil {
			queue = append(queue, *node.Right)
		}
		queue = queue[1:]
	}
	return sum
}

// HasPathSum returns true if any root-to-leaf paths equal the given sum value.
func (n Node) HasPathSum(sum int) bool {
	return false
}