package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type Linkedlist struct {
	head *Node
	tail *Node
}

func (l *Linkedlist) AddNodeAtEnd(value int) {

	newNode := &Node{data: value}

	if l.head == nil {
		l.head = newNode
	} else {
		l.tail.next = newNode
	}
	l.tail = newNode

}

func (l *Linkedlist) Print() {
	if l.head == nil {
		fmt.Println("empty")

		return
	}
	curr := l.head
	for curr != nil {
		fmt.Printf("%d", curr.data)
		curr = curr.next
	}
}

func main() {
	l := Linkedlist{}

	l.AddNodeAtEnd(1)
	l.AddNodeAtEnd(2)
	l.AddNodeAtEnd(3)
	l.AddNodeAtEnd(4)

	l.Print()
	l.InsertAfter(2,202)
	// l.DeleteByValue(4)
	fmt.Println("\nafter Deleting value ")
	l.Print()
	// fmt.Print(l.head, "\n",l.tail)
}

func (l *Linkedlist)InsertAfter(nextTo,value int){

	newNode:=&Node{}

	current:=l.head
	for current!= nil && current.data != nextTo{
		current=current.next

	}
	if current== nil{
		return
	}
	if current == l.tail{
		l.tail.next= newNode
		l.tail=newNode
		return
	}
	newNode.data=value
	newNode.next=current.next
	current.next= newNode


}


func (l *Linkedlist) DeleteByValue(value int) {

	current:=l.head

	if current != nil && current.data == value{
		current= current.next
		return
	}
	prev:=current

	for current !=nil && current.data != value{
		prev = current
		current= current.next
	}
	if current == l.tail{
		l.tail = prev
		l.tail.next= nil
		return
	}
	if current==nil{
		fmt.Println("value not in this Linkedlist")
		return
	}

	prev.next=current.next
}
