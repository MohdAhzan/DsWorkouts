package main

import "fmt"

func main(){
  
  h:=&Heap{}

  //
  //   arr:=[]int{7,3,5,9,12}
  //
  //
  //
  // for i:=range arr{
  //
  //   h.Push(arr[i])
  // }
  //
  // fmt.Println(h.elements) 
  // // h.Push(2)
  // fmt.Println(h.elements) 
  // h.Pop()
  // fmt.Println(h.elements)
  // h.Pop()
  // fmt.Println(h.elements)
  // h.Pop()
  // fmt.Println(h.elements)
  // h.Pop()
  // fmt.Println(h.elements)
  // h.Pop()
  // fmt.Println(h.elements)
  // h.Pop()
  // fmt.Println(h.elements)

   arr:=[]int{7,3,4,2,3,5,9,12,33,22,11,66}

  h.Heapify(arr)
fmt.Println(h.elements) 
  res:=[]int{}
  
  for{

    temp:=h.Pop() 
    if temp==-1{
      fmt.Println("theernnu")
      break
    }
    res=append(res,temp)
  }
  

  fmt.Println(res)
  
}

type Heap struct{
  elements []int
}

func (h *Heap)Push(val int){
  h.elements=append(h.elements, val) 
  i:=len(h.elements)-1

  for h.elements[i]< h.elements[(i-1)/2]{

      h.elements[i],h.elements[(i-1)/2]=h.elements[(i-1)/2],h.elements[i]
      
        i=(i-1)/2

  }

}

func (h *Heap)Pop()int{

  if len(h.elements)==0{
      fmt.Println("empty HEap")
    return -1
  }
    if len(h.elements)==1{
      
      res:=h.elements[0]
        h.elements=h.elements[:len(h.elements)-1]
      return res

    }
  
    res:=h.elements[0] 
    i:=0    
    h.elements[0]=h.elements[len(h.elements)-1]
    h.elements=h.elements[:len(h.elements)-1]
  
    for 2*i+1 <len(h.elements){
  
        l:=2*i+1
      r:=2*i+2

      if r<len(h.elements)&&h.elements[r]<h.elements[l]&&h.elements[i]>h.elements[r]{

        h.elements[i],h.elements[r]=h.elements[r],h.elements[i]
        i= r
      }else if h.elements[i]>h.elements[l]{
        h.elements[i],h.elements[l]=h.elements[l],h.elements[i]
          i=l

      }else{
        break
      }
        

    }

      return res 
  
}

func (h *Heap)Peek()int{
  
    return h.elements[0]

}

func (h *Heap)Heapify(arr []int){
   
    h.elements=arr

  curr:=len(h.elements)/2

  for curr>=0{
  
      i:=curr

      for 2*i+1 <len(h.elements){
        r:=2*i+2 
        l:=2*i+1

      if r < len(h.elements) && h.elements[r]< h.elements[l] && h.elements[i]> h.elements[r]{
      
            h.elements[i],h.elements[r]= h.elements[r]  , h.elements[i]  
            
          i= r
      }else if h.elements[i] > h.elements[l]{
            h.elements[i],h.elements[l]= h.elements[l]  , h.elements[i]  
             i= l

      }else {
        break
      }
    }

    curr=curr-1

  }

}
