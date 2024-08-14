package tree

import (
  "container/list"
  "fmt"

)


func (b *BST)BFS(){

  if b.Root==nil{
    return
  }

  queue:=list.New()
 queue.PushBack(b.Root)

  for queue.Len()>0{
      
    levelLength:=queue.Len()
    
    for i:=0;i<levelLength;i++{
  
      curr:= queue.Front().Value.(*Node)     
      queue.Remove(queue.Front())
      fmt.Print(curr.data," ")

      if curr.lChild!=nil{
        queue.PushBack(curr.lChild)
      }
      if curr.rChild!=nil{
        queue.PushBack(curr.rChild)
      }

    }
    fmt.Println()

  }
}
