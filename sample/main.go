package main

import "fmt"

func main() {

	link := &Linkedlist{}

	link.Add(1)
	link.Add(2)
	link.Add(3)
	link.Add(22)

	link.Print()
}

type Node struct {
	data int
	next *Node
}

type Linkedlist struct {
	head *Node
	tail *Node
}

func (l *Linkedlist) Add(value int) {

	newNode := &Node{data: value}

	if l.head == nil {
		l.head = newNode
		return
	}
	if l.head != nil {

		l.head.next = newNode
		l.tail = newNode

	}

}

func (l *Linkedlist) Print() {

	if l.head == nil {
		fmt.Println("no value")
	}
	for l.head != nil {
		fmt.Println(l.head.data)
		l.head = l.head.next
	}
}
