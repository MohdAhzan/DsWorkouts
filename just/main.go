package main

import (
	"fmt"
)

func main() {

	d := &Dlist{}

	d.Addlist(1)
	d.Addlist(2)
	d.Addlist(3)
	d.Addlist(6)
	d.Addlist(4)
	d.Addlist(5)

	d.PrintDlist()

	d.ReverseDlist()

	d.PrintDlist()
	d.Delete(5)
	d.PrintDlist()

	// d.PrintRecursively(d.head)

	// d.PrintRevereseDlist()
	//
	// d.Delete(1)
	// d.PrintDlist()
	// d.Delete(3)
	// d.PrintDlist()
	//
	// d.Delete(5)
	// d.PrintDlist()
	//
	// d.Delete(2)
	// d.PrintDlist()
	// d.Delete(4)
	// d.PrintDlist()
	// d.Delete(6)
	// d.PrintDlist()
	//
	// l := &linkedlist{}
	// l.Add(1)
	// l.Add(9)
	// l.Add(5)
	// l.Add(3)
	// l.Add(2)
	//
	// l.print()
	//
	// l.ReverseList()
	// l.print()
}

func (l *Dlist)ReverseDlist(){
	
	var prev,next *Node

	curr:=l.head

	for curr!=nil{
	
		next=curr.next
		curr.next=prev
		curr.prev=next
		prev=curr
		curr=next	

	}

	l.head=prev


}

type Dlist struct {
	head *Node
}

type Node struct {
	prev *Node
	data int
	next *Node
}

func (l *Dlist) Addlist(value int) {

	newNode := &Node{data: value}

	if l.head == nil {
		l.head = newNode
		return
	}

	curr := l.head

	for curr.next != nil {
		curr = curr.next
	}

	newNode.prev = curr
	curr.next = newNode

}

func (l *Dlist) PrintDlist() {

	if l.head == nil {

		fmt.Println("empty data in list")
		return

	}

	curr := l.head

	for curr != nil {
		fmt.Print(curr.data, "  ")
		curr = curr.next
	}

	fmt.Println()

}

func (l *Dlist) Delete(target int) {

	if l.head == nil {
		fmt.Println("no data in list")
		return
	} else if l.head.data == target {
		if l.head.next == nil {
			l.head = nil
			return
		} else {
			l.head = l.head.next
			l.head.prev = nil
			return
		}
	}

	curr := l.head

	for curr != nil && curr.data != target {

		curr = curr.next
	}

	if curr == nil {
		fmt.Println("no targettt found")
		return
	}
	if curr.next == nil && curr.data == target {

		curr.prev.next = nil
		curr = nil
		return
	}
	if curr.data == target {

		curr.prev.next = curr.next
		curr.next.prev = curr.prev
		curr = nil
	}
}

func (l *Dlist) PrintRevereseDlist() {

	if l.head == nil {

		fmt.Println("empty data in list")
		return

	}

	curr := l.head

	for curr.next != nil {
		curr = curr.next
	}

	for curr != nil {
		fmt.Print(curr.data, "  ")
		curr = curr.prev
	}

	fmt.Println()
}

func (l Dlist) PrintRecursively(node *Node) {

	if node == nil || l.head == nil {
		return
	}

	fmt.Println(node.data)

	l.PrintRecursively(node.next)
}

func replace(s string, pos int, value byte) string {

	return s[:pos] + string(value) + s[pos+1:]
}

func convert(s string, pos int, value byte) string {

	b := []byte(s)

	for i := 0; i < len(b); i++ {
		if i == pos {
			b[i] = value
		}
	}
	return string(b)
}

func reverse(s string, n int) string {

	if n < 0 {
		return ""
	}

	return string(s[n]) + reverse(s, n-1)

}

func factorial(n int) int {

	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

type node struct {
	data int
	next *node
}
type linkedlist struct {
	head *node
}

func (l *linkedlist) Add(value int) {

	newNode := &node{data: value}

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
		return
	}
}

func (l linkedlist) print() {
	if l.head == nil {
		fmt.Println("empty list")
		return
	}
	curr := l.head

	for curr != nil {
		fmt.Println(curr.data)
		curr = curr.next
	}
	fmt.Println()
}

func (l *linkedlist) delete(target int) {

	if l.head == nil {
		fmt.Println("no value ")
		return
	} else if l.head.data == target && l.head.next == nil {

		l.head = nil
		return
	} else if l.head.next != nil && l.head.data == target {
		l.head = l.head.next
		return
	}

	curr := l.head
	prev := curr
	for curr != nil && curr.data != target {

		prev = curr
		curr = curr.next
	}

	if curr == nil {
		fmt.Println("no target found", prev)
		return
	} else if curr.data == target && curr.next == nil {
		prev.next = nil
		return
	}

	prev.next = curr.next

}

func (l *linkedlist) ReverseList() {

	var prev, next *node

	curr := l.head

	for curr != nil {

		next = curr.next
		curr.next = prev
		prev = curr
		curr = next

	}

	l.head = prev
}
