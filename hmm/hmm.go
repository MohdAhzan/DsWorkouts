package main

import (
	"fmt"
)

func main() {

	l := &LinkedList{}

	l.AddintoList(11)
	l.AddintoList(22)
	l.AddintoList(33)
	l.AddintoList(44)
	l.AddintoList(55)

	l.Print()
	l.InsertAfter(55,2222333)
	l.Print()
}

 

func AddString(position int,input,str string)string{

	for i:=range str{
		if i+1==position{
			return str[:position]+input+str[position:]
		}
	}
	return "blah"
}

// func (l *LinkedList) Reverse(){}

func (l *LinkedList)Print(){

	for current := l.head; current != nil; current = current.next {
		
		fmt.Printf("%d,", current.val)
	}
	fmt.Println()	
}

type node struct {
	val  int
	next *node
}

type LinkedList struct {
	head *node
}

func (l *LinkedList) AddintoList(value int) {

	newNode := &node{val: value}

	if l.head == nil {
		l.head = newNode
		newNode.next = nil
		return
	}

	curr := l.head

	for curr.next != nil {

		curr=curr.next

	}

curr.next= newNode
newNode.next=nil

}


func (l *LinkedList)Delete(value int){

	if l.head ==nil{
		fmt.Println("empty list")
		return
	}else if l.head.val == value{

	l.head=l.head.next	

	return
	}
	current:=l.head
	prev:=l.head
	for current!=nil && current.val != value{
		prev= current	
		current=current.next

	}

	if current == nil{
		fmt.Println("no value")
		return
	}

	if current.next == nil&& current.val==value{
		current= current.next
		return
	}

	prev.next=current.next



}


func (l *LinkedList)InsertAfter(pos,value int){

	node:=&node{val: value}

	if l.head.val== pos{

		node.next=l.head.next
		l.head.next=node
		return
	}

	current:=l.head
	for current!=nil && current.val!=pos{
		current= current.next		
	}
	
	if current==nil{
	fmt.Println("position not found ")
	return
	}
	node.next=current.next
current.next=node	




}


// func (l *LinkedList)Inser
