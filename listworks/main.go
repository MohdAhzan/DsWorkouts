package main

import "fmt"


func main(){

	l:=&List{}

	l.Add(10)
	l.Add(20)
	l.Add(30)
	l.Add(40)

	l.Print()
l.Delete(20)

	l.Delete(40)
l.Delete(30)

l.Print()
}

type Node struct{
	data int
	next *Node
}

type List struct{
	head *Node
}

func (l *List)Add(value int){

	newNode:=&Node{data: value}

	if l.head == nil{

		l.head = newNode
		l.head.next = nil
		return
	}

	curr:=l.head
	for curr.next != nil{
		curr=curr.next
	}
	
	curr.next=newNode
	// newNode.next=nil



}

func (l *List)Print(){

	if l.head== nil{
		fmt.Println("empty")	
		return
	}
	curr:=l.head
	for curr!= nil{
		fmt.Println(curr.data)
		curr=curr.next
	}
	fmt.Println()
}

func (l *List)Delete(target int){


	if l.head== nil{
		fmt.Println("empty")	
		return
	}else if l.head.data==target{
	
		l.head=	l.head.next
		return

	}

	curr:=l.head
	prev:=l.head

	for curr!= nil&&curr.data!=target{

		prev=curr
		curr=curr.next
	}

	if curr == nil{

		fmt.Println("no target found")	
		return		
	}
	
	prev.next=curr.next



}
