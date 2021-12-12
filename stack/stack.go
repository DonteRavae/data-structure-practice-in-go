/**
* Written by Donte Littlejohn
**/
package stack

import (
	"errors"
	"fmt"
)

type node struct {
	data int
	next *node
	previous *node
}

type Stack struct {
	head *node
	tail *node
	length int
}

func (s *Stack) Push(dataToAdd int) {
	nodeToAdd := node{
		data: dataToAdd,
	}
	if s.head == nil {
		s.head = &nodeToAdd
		s.tail = s.head
		s.length++
		return
	}
	nodeToAdd.previous = s.tail
	s.tail.next = &nodeToAdd
	s.tail = &nodeToAdd
	s.length++
}

func (s *Stack) Pop() (*node, error) {
	if s.head == nil {
		return nil, errors.New("empty stack")
	}
	nodeToDelete := s.tail
	s.tail = nodeToDelete.previous
	s.tail.next = nil
	s.length--
	return nodeToDelete, nil
}

func (s Stack) Length() int {
	return s.length
}

func (s Stack) PrintStackFromTop() {
	if s.tail == nil {
		fmt.Println("Empty stack")
		return
	}
	tmp := s.tail
	for tmp != nil {
		fmt.Printf("%d", tmp.data)
		tmp = tmp.previous
		fmt.Println()
	}
	fmt.Println()
}

func (s Stack) PrintStackFromBottom() {
	if s.head == nil {
		fmt.Println("Empty stack")
		return
	}
	tmp := s.head
	for tmp != nil {
		fmt.Printf("%d", tmp.data)
		tmp = tmp.next
		fmt.Println()
	}
	fmt.Println()
}

func (s Stack) Peak() *node {
	if s.tail == nil {
		return nil
	}
	return s.tail
}