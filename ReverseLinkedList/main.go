package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func main() {

	l := &LinkedList{}
	l.Add(1)
	l.Add(12)
	l.Add(33)
	l.Add(123)
	l.Add(1234)
	l.Print()
	// l.ReverseList()
	l.Delete(1234)
	l.Delete(1)
	l.Delete(33)
	l.Print()
}
func (l *LinkedList) Add(value int) {

	newNode := &Node{data: value}

	if l.head == nil {
		l.head = newNode
		return
	}

	curr := l.head

	for curr.next != nil {

		curr = curr.next

	}

	if curr.next == nil {
		curr.next = newNode
	}

}

func (l *LinkedList) Print() {

	curr := l.head

	for curr != nil {
		fmt.Println(curr.data)
		curr = curr.next
	}

	fmt.Println()

}

func (l *LinkedList) ReverseList() {

	var prev, next *Node

	curr := l.head

	for curr != nil {

		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}

	l.head = prev

}

func (l *LinkedList) Delete(target int) {

	if l.head == nil {

		fmt.Println("empty list")
		return
	} else if l.head.data == target {

		l.head = l.head.next
		return
	}

	curr := l.head

	for curr.next != nil && curr.next.data != target {
		curr = curr.next

		if curr.next.data == target && curr.next.next == nil {
			fmt.Println("deleting at last node...")
			curr.next = nil
			return
		}
		if curr.next == nil {
			fmt.Println("target no found")
			return
		}

		if curr.next.data == target {

			curr.next = curr.next.next
			return
		}

	}
}
