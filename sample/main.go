package main

import "fmt"



type Node struct {
	prev *Node
	val  int
	next *Node
}

type DLinkedList struct {
	head *Node
	tail *Node
}

func (l *DLinkedList) Add(value int) {

	newNode := &Node{val: value}

	if l.head == nil {

		l.head = newNode
		l.tail = newNode
	} else {

		l.tail.next = newNode
		newNode.prev = l.tail
		l.tail = newNode

	}

}

func (new *DLinkedList) Print() {
	if new.head==nil{
		fmt.Println("empty list")
		return
	}


	for node := new.head; node != nil; node = node.next {

		fmt.Println(node.val)
	}
	fmt.Println()
	fmt.Println()
}
func main() {

	new := &DLinkedList{}

	new.Add(2)
	new.Add(3)
	new.Add(4)
	new.Add(7)
	new.Print()
	// for node := new.head; node != nil; node = node.next {
	// 	fmt.Println(node.val)
	// }

	new.Deletebtwn(4)
	new.Print()
	new.Deletebtwn(4)
	new.Print()
	new.Deletebtwn(3)
	new.Print()
	new.Deletebtwn(2)
	new.Print()
	new.Deletebtwn(7)
	new.Print()

	// fmt.Println(new.tail.val,"was this deleted?")

}

func (l *DLinkedList) Deletebtwn(value int) {

	if l.head == nil {
		fmt.Println("empty linked list")
		return
	} else if l.head.val == value {
		if l.head.next == nil{
			l.head=nil
			return
		}

		l.head = l.head.next
		l.head.prev = nil

		return
	}

	current := l.head
	for current != nil && current.val != value {

		current = current.next
	}

	if current == nil {
		fmt.Println("valur is not here  ")
		return

	}
	if current.val == l.tail.val {

		current.prev.next = nil
		l.tail = current.prev
		return
	}
	current.prev.next = current.next
	current.next.prev = current.prev

}

