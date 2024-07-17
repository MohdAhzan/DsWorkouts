package main

import (
	"fmt"
)

func main(){ 

	// l := &linkedlist{}
	//
	// l.AddintoList(11)
	// l.AddintoList(22)
	// l.AddintoList(33)
	// l.AddintoList(44)
	// l.AddintoList(55)
	//
	// l.print()
	// l.delete(11)
	// l.print()
	// l.delete(44)
	// l.print()
	// l.delete(55)
	// l.print()
	// l.delete(3)
	// l.print()
	// l.delete(33)
	// l.print()
	//
	// l.delete(22)
	// l.print()
	// l.print()

	st:="vazheeee"

	fmt.Println(stringy(st,len(st)-1))
}

func stringy(s string,n int)string{
	if n==0{
		return ""
	}

	return string(s[n])+stringy(s,n-1)
}

type node struct{
	data int
	next *node
}
type linkedlist struct{
	head *node
}
func (l *linkedlist)AddintoList(val int){

	
	newNode:=&node{data : val}

	if l.head==nil{
		l.head=newNode
		return
	}

	curr:=l.head

	for curr.next!=nil{

		curr=curr.next
	}

		curr.next=newNode
		return
}

func (l linkedlist)print(){
	if l.head==nil{
		fmt.Println("empty.....")
		return
	}
	curr:=l.head

	for curr!=nil{
		fmt.Print(curr.data,"  ")	
		curr=curr.next
	}
	
	fmt.Println()

}

func (l *linkedlist)delete(target int){

	if l.head==nil{
		fmt.Println("no value")
		return
	}else if l.head.data==target&&l.head.next!=nil{
		
	l.head=l.head.next
	return
	}else if l.head.data==target{
		l.head=nil
		return
	}
	curr:=l.head
	for curr.next!=nil&&curr.next.data!=target{
		curr=curr.next
	}
	
	if curr.next==nil{
			fmt.Println("target not in list")
			return
	}
	if curr.next.next==nil&&curr.next.data==target{
		curr.next=nil
		return
	}

	curr.next=curr.next.next

}

func (l *linkedlist)insertAfter(after,value int){

	
newNode:=&node{data: value}

if l.head==nil{
	l.head=newNode
	return
}else if l.head.data==after&&l.head.next!=nil{
	newNode.next=l.head.next
	l.head.next=newNode
	return
}else if l.head.next==nil{
	l.head.next=newNode
	return
}

curr:=l.head

for curr!=nil&&curr.data!=after{
	curr=curr.next
}

if curr==nil{
	fmt.Println("after value not found ")
	return
}else if curr.data==after &&curr.next==nil{
	curr.next=newNode
	return
}
newNode.next=curr.next
curr.next=newNode

}





















