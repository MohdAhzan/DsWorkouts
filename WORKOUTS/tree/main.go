package main

import "fmt"

func  main(){

  t:=&BST{}

    
    t.InsertBST(7)
    t.InsertBST(44)
    t.InsertBST(29)
    t.InsertBST(3)
    t.InsertBST(4)
    t.InsertBST(1)

  t.PrintInOrder(t.root)
  fmt.Println("")
  t.PrintPreOrder(t.root)
    fmt.Println("")
  t.PrintPostOrder(t.root)



}

type treeNode struct{
  lnode *treeNode
  data int
  rnode *treeNode
}

type BST struct{
  root *treeNode
}

func (b *BST)InsertBST(value int){

  newNode:=&treeNode{data: value}
  if b.root==nil{
    b.root=newNode
    return 
  }
  curr:=b.root
  for curr!=nil{

    if value < curr.data{
      if curr.lnode==nil{
        curr.lnode=newNode
        return
      }else{
        curr=curr.lnode
      }

    }else if value > curr.data{

      if curr.rnode==nil{
        curr.rnode=newNode
        return
      }else{
          curr=curr.rnode
      }
    }
  }


}


func (b *BST)PrintInOrder(curr *treeNode){
  
  if curr==nil{
    return 
  }
  b.PrintInOrder(curr.lnode)
  fmt.Println(">",curr.data)
  b.PrintInOrder(curr.rnode)
}

func (b *BST)PrintPostOrder(curr *treeNode){
  
  if curr==nil{
    return 
  }
  b.PrintInOrder(curr.lnode)
  b.PrintInOrder(curr.rnode)
  fmt.Println(">",curr.data)
}

func (b *BST)PrintPreOrder(curr *treeNode){
  
  if curr==nil{
    return 
  }
  fmt.Println(">",curr.data)
  b.PrintInOrder(curr.lnode)
  b.PrintInOrder(curr.rnode)
}


