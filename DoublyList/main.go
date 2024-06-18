package main

import "fmt"

type Node struct {
	prev *Node
	data int
	next *Node
}

type DLinkedList struct {
	head *Node
	tail *Node
}

func main() {

	l := &DLinkedList{}
	l.Print()
	l.Add(11)
	l.Add(22)
	l.Add(55)
	l.Add(33)
	l.Add(44)
	// l.DeleteByValue(11)
	// l.DeleteByValue(44)
	l.Print()

}

func (l *DLinkedList) Add(value int) {

	newNode := &Node{data: value}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		newNode.prev = l.tail
		l.tail = newNode

	}

}

func (l *DLinkedList) Print() {
	curr := l.head
	if curr == nil {
		fmt.Println("empty dlinked list")
		return
	}
	for curr != nil {
		fmt.Println(curr.data)
		curr = curr.next
	}

}

func (l *DLinkedList) DeleteByValue(value int) {

	current := l.head

	if current != nil && current.data == value {
		l.head = current.next
		current.next.prev = nil
		return
	}

	for current != nil && current.data != value {

		current = current.next

	}

	if current == l.tail {
		fmt.Println("tail", current.prev.next)
		current.prev.next = nil
		return

	}
	fmt.Println(current, current.prev)
	current.prev.next = current.next
	current.next.prev = current.prev
}
