package main

import "fmt"

// Stack represents stack that holds a slice
type Stack struct {
	items []int
}

// Queue represents a queue that holds a slice
type Queue struct {
	items []int
}

func main() {
	//STACK EXAMPLE
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)


	//QUEUE EXAMPLE
	myQueue := Queue{}
	myQueue.Enqueue(33)
	myQueue.Enqueue(72)
	myQueue.Enqueue(48)
	fmt.Println(myQueue)
	myQueue.Dequeue()
	fmt.Println(myQueue)
}

// Push will add a value to the end
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop will remove value at the end and RETURNS the removed value
func (s *Stack) Pop() int {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

// Enqueue adds a value to the end of the list
func (q *Queue) Enqueue (i int) {
	q.items = append(q.items, i)
}

// Dequeue removes a value from the beginning of the list and RETURNS the removed value
func (q *Queue) Dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}