package main

import (
  "fmt"
)

func main(){
  
  l:=&List{}
  l.Add(123)
  l.Add(33)
  l.Add(63)
  l.Add(88)

  l.PrintList()

    
}

type Node struct{
   
  data int
  next *Node

}

type List struct{
  
  head *Node
  tail *Node
}

func (l *List)Add(value int){

    newNode:=&Node{data: value}

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

func (l List)PrintList(){
  
  curr:=l.head

  for curr!=nil{
    fmt.Print(">",curr.data)
    curr=curr.next
  }

  fmt.Println()
  
}
