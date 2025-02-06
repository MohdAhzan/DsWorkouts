package main

import (
	"fmt"
)

func main(){

  l:=&LINK{}


    l.Head = &node{data: 1}
    l.Head.next = &node{data: 2}
    l.Head.next.next = &node{data: 3}
    // l.Head.next.next.next = l.Head 

    l.CHECK() // Output: "CIRCULAR"
  // l.AddHM(1)
  // l.AddHM(2)
  // l.AddHM(3)
  // l.AddHM(4)
  // l.AddHM(5)
  
  // l.Print()
  // nn:=l.GetNthNOde(4)
  // if nn!=nil{
  //   fmt.Println("data",nn.data,"  node",nn)
  // }

}

func (l *LINK)CHECK(){
  
  if l.Head==nil{
    return
  }
  first:=l.Head
  second:=l.Head.next

  for second!=nil&&second.next!=nil{
  
      if first==second{
      fmt.Println("CIRCULAR ")
        return
    }
  
    first=first.next
    second=second.next.next
  }

  fmt.Println("not circular single linked list")
}


type LINK struct{
  Head *node
}

func (l *LINK)GetNthNOde(n int)*node{
  if l.Head==nil{
    return nil
  }
  curr:=l.Head
  for i:=1;i<=n;i++{
    if curr==nil {
      fmt.Printf("length of linked is only %d not %d enter valid range",i,n)
      return nil
    }
    if i==n{
      return curr 
    }else{
      curr=curr.next
    }
  }
  return  nil

}

type node struct{
  data int
  next *node
}

func (l *LINK)Delete(target int){
  if l.Head==nil{
    return
  }else if l.Head.data==target{
    l.Head=l.Head.next
  }
  curr:=l.Head
  for curr.next!=nil&&curr.next.data!=target{
    curr=curr.next
  }
  if curr.next!=nil{
    curr.next=curr.next.next
  }else{
    curr.next=nil
  }
}

func (l *LINK)ReverseREcursion(){
  l.Head=reverse(l.Head)
}
func reverse(node *node) *node{

  if node==nil || node.next == nil{
    return node
  }
  newNode:=reverse(node.next)

  node.next.next=node
  node.next=nil

  return newNode
}

func (l *LINK)Reverse(){

  if l.Head==nil{
    fmt.Println("empty")
    return
  }
  var prev,next *node
  curr:=l.Head
  for curr!=nil{
    next=curr.next
    curr.next=prev
    prev=curr
    curr=next
  }
  l.Head=prev
}

func (l *LINK)AddHM(value int){

  newNode:=&node{data: value}

  if l.Head==nil{
    l.Head=newNode
    return
  }
  curr:=l.Head
  for curr.next!=nil{
    curr=curr.next
  }
  curr.next=newNode
}

func (l *LINK)Print(){
  curr:=l.Head
  for curr!=nil{
    fmt.Print(" ",curr.data)
    curr=curr.next
  }
  fmt.Println()
}


