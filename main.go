/**
* Written by Donte Littlejohn
**/
package main

import (
	"fmt"
	"github.com/donteravae/data-structure-practice/linked-list"
	"github.com/donteravae/data-structure-practice/binary-tree"
	"github.com/donteravae/data-structure-practice/stack"
	"github.com/donteravae/data-structure-practice/queue"
)

func main() {
	fmt.Println("Data Structures and Algorithms in Golang")
	fmt.Println()
	
	// Linked Lists
	fmt.Println("Linked List Example")
	testList := linkedList.LinkedList{}
	testList.Add(3)
	testList.Add(2)
	testList.Add(24)
	testList.Add(8)
	testList.Add(89)
	fmt.Printf("The length of the list is %d\n", testList.Length())
	testList.PrintListData()
	testList.Remove(24)
	testList.PrintListData()
	testList.Remove(89)
	testList.PrintListData()
	testList.Remove(3)
	fmt.Printf("The length of the list is %d\n", testList.Length())
	testList.PrintListData()
	fmt.Println()

	// Stacks
	fmt.Println("Stack Example")
	myStack := stack.Stack{}
	myStack.PrintStackFromTop()
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
	myStack.PrintStackFromBottom()
	popped, err := myStack.Pop()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(popped)
	myStack.PrintStackFromTop()

	// Queues
	fmt.Println("Queue Example")
	myQueue := queue.Queue{}
	myQueue.Enqueue(33)
	myQueue.Enqueue(72)
	myQueue.Enqueue(48)
	myQueue.Print()
	myQueue.Dequeue()
	myQueue.Print()

	// Binary Tree
	fmt.Println("Binary Tree Example")
	tree := binaryTree.Node{Data: 89}
	tree.Insert(22)
	tree.Insert(190)
	tree.Insert(54)
	tree.Insert(647)
	tree.Insert(831)
	tree.Insert(656)
	// Print Tree
	// fmt.Println(tree)
	// fmt.Println(tree.Right.Left)
	// fmt.Println(tree.Left)
	// fmt.Println()
}