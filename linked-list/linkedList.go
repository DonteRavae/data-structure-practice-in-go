package linkedList

import "fmt"

type node struct {
	data int
	next *node
}

type LinkedList struct {
	head *node
	length int
}

func (ll LinkedList) Length() int {
	return ll.length
}

func (ll *LinkedList) Add(dataToAdd int) {
	nodeToAdd := &node{ data: dataToAdd }
	if ll.head == nil {
		ll.head = nodeToAdd
		ll.length++
		return
	}

	tmp := ll.head
	for tmp.next != nil {
		tmp = tmp.next
	}
	tmp.next = nodeToAdd
	ll.length++
}

func (ll LinkedList) PrintListData() {
	toPrint := ll.head
	for toPrint != nil {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
	}
	fmt.Println()
}

func (ll *LinkedList) Remove(val int) {
	if ll.head == nil {
		return
	}

	if ll.head.data == val {
		ll.head = ll.head.next
		ll.length--
		return
	}

	previousToDelete := ll.head
	for previousToDelete.next.data != val {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	ll.length--
}