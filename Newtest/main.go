package main

import (
  "fmt"

)

func main(){

  b:=&BST{}

  arr:=[]int{34,24,12,4,66,3,62,33}

  for i:=range arr{

    b.Insert(arr[i])

  }

  PrintTree(b.root)

}

type BST struct{

  root *node
}
type node struct{

  left *node
  data int
  right *node
}


func (b *BST)Insert(value int){

  newNode:=&node{data: value} 

  if b.root==nil{
    b.root=newNode
    return
  }
  curr:=b.root

  for curr!=nil{

    if value < curr.data{

      if curr.left==nil{
        curr.left=newNode
        return

      }else{
        curr=curr.left
      }
    }else if  value > curr.data{
     
      if curr.right==nil{
        curr.right=newNode
        return

      }else{
        curr=curr.right
      }

    }

  }

}

func PrintTree(curr *node){

  if curr==nil{
    return
  }
  PrintTree(curr.left)
  fmt.Print(curr.data," ")
  PrintTree(curr.right)
}
