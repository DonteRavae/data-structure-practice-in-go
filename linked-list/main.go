package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
	length int
}

func main() {
	testList := linkedList{}
	node1 := node{data: 3}
	node2 := node{data: 2}
	node3 := node{data: 24}
	node4 := node{data: 8}
	node5 := node{data: 21}

	testList.prepend(&node1)
	testList.prepend(&node2)
	testList.prepend(&node3)
	testList.prepend(&node4)
	testList.prepend(&node5)

	testList.printListData()
	testList.removeData(24)
	testList.printListData()
	testList.removeData(89)
	testList.printListData()
	testList.removeData(3)
	testList.printListData()
}

func (l *linkedList) prepend(n *node) {
	previousHead := l.head
	l.head = n
	l.head.next = previousHead
	l.length++
}

func (l linkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Println()
}

func (l *linkedList) removeData(val int) {
	if l.length == 0 {
		return
	}

	if l.head.data == val {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.data != val {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}