package main

import (
	"fmt"
)

func main() {
   
  h:=&Heap{}
  h.Push(33)
  h.Push(22)
  h.Push(11)
  h.Push(55)
  h.Push(88)
  h.Push(44)
  fmt.Println(h.heap) 

  h.Pop()
  fmt.Println(h.heap) 
  h.Pop()
  fmt.Println(h.heap) 
  h.Pop()
  fmt.Println(h.heap) 
  h.Pop()
  fmt.Println(h.heap) 
  h.Pop()
  fmt.Println(h.heap) 
  h.Pop()
  fmt.Println(h.heap) 
  h.Pop()
  fmt.Println(h.heap) 
  h.Pop()
  fmt.Println(h.heap) 
}

type Heap struct{

  heap []int
}

func (h *Heap)Push(val int){


  h.heap = append(h.heap, val)
  i:=len(h.heap)-1
  for h.heap[i] < h.heap[(i-1)/2]{
    h.heap[i],h.heap[(i-1)/2]=h.heap[(i-1)/2],h.heap[i]
    i=(i-1)
  }
}

func (h *Heap)Pop()int{
    
    if len(h.heap)==0{
    fmt.Println("empty heap") 
      return -999
  }

  if len(h.heap)==1{

    poped:=h.heap[0]
    h.heap=h.heap[:0]
    return poped
  }
  poped:=h.heap[0] 
  h.heap[0]=h.heap[len(h.heap)-1]
  h.heap=h.heap[:len(h.heap)-1]

  i:=0

 
  
  for 2*i+1 < len(h.heap){
  
      r:=2*i+2
      l:=2*i+1

      if r < len(h.heap) && h.heap[r] < h.heap[l] && h.heap[i] > h.heap[r]{

      h.heap[i],h.heap[r]=h.heap[r],h.heap[i]
      i=r
   }else if  h.heap[l]< h.heap[i]{
    h.heap[i],h.heap[l]=h.heap[l],h.heap[i]
      i=l

    }else{
      break
    }

  }

  return poped


}
