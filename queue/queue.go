/**
* Written by Donte Littlejohn
**/
package queue

import "fmt"

type node struct {
	data int
	next *node
}

type Queue struct {
	head   *node
	tail   *node
	length int
}

func (q *Queue) Enqueue(dataToAdd int) {
	nodeToAdd := node{
		data: dataToAdd,
	}
	if q.head == nil {
		q.head = &nodeToAdd
		q.tail = q.head
		q.length++
		return
	}
	q.tail.next = &nodeToAdd
	q.tail = &nodeToAdd
	q.length++
}

func (q *Queue) Dequeue() *node {
	if q.head == nil {
		fmt.Println("empty queue")
		return nil
	}
	nodeToRemove := q.head
	q.head = q.head.next
	q.length--
	return nodeToRemove
}

func (q Queue) Peak() *node {
	if q.head == nil {
		return nil
	}
	return q.head
}

func (q Queue) Print() {
	if q.head == nil {
		fmt.Println("empty queue")
		return
	}
	toPrint := q.head
	for toPrint != nil {
		fmt.Printf("%d", toPrint.data)
		toPrint = toPrint.next
		fmt.Println()
	}
	fmt.Println()
}