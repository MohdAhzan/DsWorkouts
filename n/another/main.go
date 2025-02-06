package main

import "fmt"

func main(){

    arr:=[]int{3,5,1,7,2}
  d:=&DList{} 
  for _,n := range arr {
    d.Add(n)
  }
  d.Print()
    d.InsertAfter(4,0) 
  d.Print()
  d.InsertAfter(5,0) 
  d.Print()
  d.InsertAfter(2,11) 
  d.Print()
  d.InsertAfter(3,666) 
  d.Print()
 
}

type DList struct{
  Head *Node
}
type Node struct{
    prev *Node
    data int
    next *Node
}

func (d *DList)Add(v int){
 
  nn:=&Node{data: v}
  if d.Head==nil{
  d.Head=nn
    return
  }
  curr:=d.Head
  for curr.next!=nil{
     curr=curr.next 
  }
  curr.next=nn
  nn.prev=curr
}

func (d *DList)Reverse(){
  
  curr:=d.Head
  var temp *Node
  for curr!=nil{
    
    curr.next,curr.prev=curr.prev,curr.next
    temp=curr
    curr=curr.prev
  }
  d.Head=temp

}

func (d *DList)Print(){
  c:=d.Head
  for c!=nil{
    fmt.Print("  ",c.data)
    c=c.next
  }
  fmt.Println("")

}

func (d *DList)Delete(t int){

    if d.Head==nil{
    return
  }
  curr:=d.Head
  if curr!=nil&&curr.data!=t{
      curr=curr.next
  }
  
  if curr!=nil{
    curr.prev.next=curr.next
    curr.next.prev=curr.prev
  }

}

func (d *DList)InsertAfter(after,value int){
 
    newNode:=&Node{data: value}
  
  if d.Head==nil{
      d.Head=newNode
      return
  }

    curr:=d.Head

  for curr!=nil&& curr.data!=after{ 
      curr=curr.next
  }

  if curr==nil{
    fmt.Println("after value not found in this list")
     return   
  }
  if curr.data==after&&curr.next!=nil{
      newNode.next=curr.next
      curr.next=newNode
      
      return
  }else{
    curr.next=newNode
  }

}
