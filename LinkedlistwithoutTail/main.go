package main

import "fmt"

func main(){
	l:=LinkedList{}

	l.Add(1)
	l.Add(2)
	l.Add(3)	
	l.Print()
fmt.Println()
	l.InsertAfter(2,888)
	l.Print()
}

type LinkedList struct{
	Head *Node
}
type Node struct{

	data int
	next *Node
}



func (l *LinkedList)InsertAfter(after,value int){
	
	newNode:= &Node{data: value}

	current:=l.Head
	
	if l.Head == nil{
		fmt.Println("Empty LinkedList")
		return
	}else if l.Head.data == after{
		newNode.next = current.next
		current.next = newNode
		return
	}

	for current!=nil && current.data != after{
		
		current= current.next

	}

	if current==nil{
		fmt.Println("after value not found in this LinkedList")
		return

	}

	newNode.next= current.next
	current.next=newNode

	
}

func (l *LinkedList)Add(value int){
	
	newNode:=&Node{data: value}

	if l.Head==nil{
		l.Head = newNode
		l.Head.next= nil
		return
	}
	current:=l.Head
	for current.next != nil{
		current= current.next
	}
	current.next= newNode
	current.next.next= nil


}

func (l *LinkedList)Print(){

	current:= l.Head
	if l.Head == nil{

		fmt.Println("Empty LinkedList")
	}
	for current != nil {
		fmt.Println(current.data)	
		current=current.next	

	}
}


func (l *LinkedList)Delete(value int){
	
	if l.Head == nil{
		fmt.Println("Empty linked list broo")
		return
	}

	if l.Head.data == value {
		l.Head=l.Head.next
		return 
	}
	current:=l.Head
	prev:=l.Head
	for current!=nil && current.data != value{

		prev=current
		current=current.next
	}


	if current== nil{
		fmt.Printf("%d is not in this linked list",value)

		return
	}
	
	prev.next = current.next
}
