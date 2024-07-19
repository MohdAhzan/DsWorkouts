package main

import (
	"fmt"
)

func main() {

	l := &list{}

	// arr := []int{1, 2, 5, 1, 2, 5, 7, 8, 1, 5, 7, 3, 3, 5, 5}
	arr:=[]int{1,1,2,2,3,3,4,4,5,5}
	for i := 0; i < len(arr); i++ {

		l.Add(arr[i])

	}
	l.print()

	mapcount := make(map[int]bool)

	curr := l.head
	prev:=curr
	for curr != nil {
		if !mapcount[curr.val] {
			mapcount[curr.val] = true
		} else {
	
			prev.next=curr.next
			prev=curr
			curr=nil
			// l.Delete(curr.val)

		}
		curr = prev.next
	}

	l.print()

}
func (l list) print() {

	if l.head == nil {
		fmt.Println("empty")
		return

	}
	curr := l.head

	for curr != nil {

		fmt.Println(curr.val)
		curr = curr.next
	}
	fmt.Println()
}

type Node struct {
	val  int
	next *Node
}

type list struct {
	head *Node
}

func (l *list) Add(value int) {

	newNode := &Node{val: value}

	if l.head == nil {
		l.head = newNode
		return
	}

	curr := l.head

	for curr.next != nil {
		curr = curr.next
	}

	curr.next = newNode

}

func (l *list) Delete(target int) {
	if l.head == nil {
		fmt.Println("empty")
		return
	} else if l.head.val == target {
		if l.head.next != nil {
			l.head = l.head.next
			return
		} else {
			l.head = nil
			return
		}
	}

	curr := l.head
	prev := curr

	for curr != nil && curr.val != target {

		prev = curr
		curr = curr.next
	}

	if curr == nil {
		fmt.Println("no target found")
		return
	} else if curr.next == nil && curr.val == target {

		prev.next = nil
		return

	}

	prev.next = curr.next
	curr = nil

}
