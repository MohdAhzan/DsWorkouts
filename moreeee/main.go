package main

import "fmt"

func main(){

	l:=&List{}
	for i:=1;i<=10;i++{


	l.Add(i)
	}
	

	// print list of 1 to 10
	// l.Print()
	// l.AddatHead(99)
	l.Print()

	l.InsertAfter(24,10)
l.Delete(1)

	l.Print()

	l.Delete(24)
l.Delete(8)
l.Delete(10)

l.Delete(3)
l.Delete(4)
l.Delete(9)
l.Delete(10)
l.Print()
}




type Node struct{
data int
next *Node
}

type List struct{
	
	head *Node

}

func (l *List)Add(val int){
	newNode:=&Node{data: val}

	if l.head==nil{
	
		l.head=newNode
		return

	}
	
	curr:=l.head
	for curr.next!=nil{
		curr=curr.next
	}
	
	curr.next=newNode
	
}

func (l *List)Print(){
	
	curr:=l.head
	for curr!=nil{
	
		fmt.Println(curr.data)
		curr=curr.next
	}
	
	fmt.Println()
}



func (l *List)AddatHead(value int){
	new:=&Node{data: value}

	if l.head==nil{
	
		l.head=new

		return
	}else{

	
		new.next=l.head

		l.head=new
		return

	}
}


func (l *List)InsertAfter(value,target int){
	
	
	newNode:=&Node{data: value}

	if l.head==nil{
	
		fmt.Println("no values in linkdlist")
		return
	}
	
	if l.head!=nil&&target==l.head.data{

	newNode.next=l.head.next
	l.head.next=newNode
	return
	}
	curr:=l.head

	for curr!=nil&&curr.data!=target{

	
		curr=curr.next
	}
	
	if curr==nil{
	
		fmt.Println("target value not found in list")
		return
	}
	if curr.data==target{
	
		newNode.next=curr.next

		curr.next=newNode
	}

}

func (l *List)Delete(target int){
	if l.head==nil{
	
		fmt.Println("no values in linkdlist")
		return
	}else if l.head.data == target{

	l.head=l.head.next
	return

	}
	
	curr:=l.head
	prev:=l.head
	
	for curr!=nil&&target!=curr.data{
		prev=curr
		curr=curr.next
	}
	if curr==nil{
	
		fmt.Println("target not found")
		return
	}else if  target==curr.data{
		
		prev.next=curr.next
	
		return
	}

}
