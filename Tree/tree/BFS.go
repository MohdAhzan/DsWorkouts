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
  level:=0
  for queue.Len()>0{
  
    fmt.Print("level: ",level,"   ")
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
    level++
    fmt.Println()

  }
}
