package main

import "fmt"

func main() {

	List := &LinkedList{}

	arr := []int{12, 32, 43, 5, 93, 34, 22, 11, 231, 9393}

	for _, num := range arr {

		List.AddFromArray(num)
	}

	List.Print()

}

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (l *LinkedList) AddFromArray(value int) {

	newnode := &Node{data: value}

	if l.head == nil {
		l.head = newnode
		l.tail = l.head
		return
	} else {
		l.tail.next = newnode
		l.tail=newnode
	
		return
	}

}

func (l *LinkedList) Print() {
current:=l.head
	for current!= nil {
		fmt.Println(current.data)
		current=current.next
	}
}
