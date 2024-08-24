package main

import (
	"container/list"
	"fmt"
)

func main(){
  
    b:=&BST{}
  b.Insert(3)
  b.Insert(8)
  b.Insert(7)
  b.Insert(6)
  b.Insert(61)
  b.Insert(16)
  
  b.BFS()
  
    
  
}




type BST struct{
  Root *Node
}
type Node struct{
  lChild *Node
  data int
  rChild *Node
}


func (b *BST)Insert(value int){

  if b.Root==nil{
    b.Root=&Node{data: value}
    return
  }

  curr:=b.Root
  newNode:=&Node{data: value}
  for curr!=nil{

    if value > curr.data {
      if curr.rChild==nil{
        curr.rChild=newNode
        return

      }else{
        curr=curr.rChild

      }

    }else if value <= curr.data {

        fmt.Println("here")

      if curr.lChild==nil{
        curr.lChild=newNode
        return

      }else{
        curr=curr.lChild
      }
    }
  }

}

func  PrintInOrder(curr *Node){

      if curr==nil{
      return
  }
    PrintInOrder(curr.lChild)
    fmt.Print(" ",curr.data)
    PrintInOrder(curr.rChild)

    
}
func  PrintPreOrder(curr *Node){

      if curr==nil{
      return
  }
    PrintPreOrder(curr.lChild)
    PrintPreOrder(curr.rChild)
  fmt.Print(" ",curr.data)

    
}

func  PrintPostOrder(curr *Node){

      if curr==nil{
      return
  }
    PrintPostOrder(curr.lChild)
    PrintPostOrder(curr.rChild)
  fmt.Print(" ",curr.data)

    
}

func (b *BST)Contains(target int)bool{
 
  if b.Root==nil{
    fmt.Println("empty BST")
    return false
  }
  
  curr:=b.Root

  for curr!=nil{
    
      if curr.data==target{
      fmt.Println("target ",curr.data,"found")
      return true
    }else if target > curr.data {
      curr=curr.rChild
    }else{
      curr= curr.lChild
    }


  }
   
  return false

}
func (b *BST)Delete(target int){
  
  if b.Root==nil{
    fmt.Println("Empty BST ")
    return
  }
  b.Root=DeleteBst(b.Root,target)
}

func DeleteBst(curr *Node,target int)*Node{
  
    if curr==nil{
    return nil
  }
    
  if target < curr.data{
  
      curr.lChild=DeleteBst(curr.lChild,target)
  }else if target > curr.data{

    curr.rChild =DeleteBst(curr.rChild,target)
  }else{

      if curr.rChild==nil{
      return curr.lChild
    }else if curr.lChild==nil{
      return curr.rChild
    }else{
  
        min:=minNode(curr.rChild)
        curr.data=min.data
        curr.rChild=DeleteBst(curr.rChild,min.data) 
        

    }

  }

  return curr

}

func minNode(curr *Node)*Node{
  
    for curr!=nil&&curr.lChild!=nil{
    curr=curr.lChild
  }

  return curr

}

func (b *BST)BFS(){


  if b.Root==nil{

    return
  }
    queue:=list.New()
    queue.PushBack(b.Root)  
  
    for queue.Len()>0{
  
      level:=queue.Len()

      for i:=0;i<level;i++{

        curr:=queue.Front().Value.(*Node)
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



/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Trie struct{

  root *TrieNode

}
type TrieNode struct{

  child map[rune]*TrieNode
  isEnd bool
}

func NewTrie()*Trie{

  return &Trie{root: &TrieNode{child: make(map[rune]*TrieNode)}}
}

func (t *Trie)Insert(word string ){

  curr:=t.root

  for _,char:=range word{

    if _,ok:=curr.child[char];!ok{

      curr.child[char]=&TrieNode{child: make(map[rune]*TrieNode)}

    }

    curr=curr.child[char]

  }

  curr.isEnd=true

}

func (t *Trie)Contains(word string)bool{

  curr:=t.root

  for _,ch:=range word{

    if _,ok:=curr.child[ch];!ok{

      return false

    }
    curr=curr.child[ch]
  }

  return true
}

func (t *Trie)PrintALLWORDS(){

  t.collectWords(t.root,"")

}
func (t *Trie)collectWords(curr *TrieNode,prefix string){

  if curr.isEnd{

    fmt.Println(prefix) 
  }

  for ch,currNode:=range curr.child{

    t.collectWords(currNode,prefix+string(ch))

  } 

}




///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////


func (t *Heap )HeapSort(arr []int){

  t.elements=arr

  t.Heapify(t.elements)



  res:=[]int{}
  for  {
    temp:=t.Pop() 

    if temp == -999{
      break
    } 

    res=append(res,temp)

  }


  copy(arr,res)

  return 
}


type Heap struct{

  elements []int
}

func (h *Heap)Push(value int){


  h.elements = append(h.elements, value)

  i:=len(h.elements)-1

  for h.elements[i]<h.elements[(i-1)/2]{

    h.elements[i],h.elements[(i-1)/2]=h.elements[(i-1)/2],h.elements[i]
    i=(i-1)/2

  }

}

func (h *Heap)Heapify(arr []int){

  h.elements=arr

  curr:=(len(h.elements)-1)/2

  for curr >=0{

    i:=curr

    for 2*i+1 <len(h.elements){

      l:=2*i+1
      r:=2*i+2

      if r< len(h.elements)&& h.elements[r]<h.elements[l] && h.elements[r]< h.elements[i]{
        h.elements[i],h.elements[r]=h.elements[r],h.elements[i]
        i=r
      }else if h.elements[l] < h.elements[i]{
        h.elements[i],h.elements[l]=h.elements[l],h.elements[i]
        i=l

      }else{
        break
      }
    }


    curr--

  } 


}

func (h *Heap)Pop( )int{

  if len(h.elements)==0{
    fmt.Println("Empty heap")
    return -999
  }

  if len(h.elements)==1{
    removed:=h.elements[0]
    h.elements=h.elements[:0]
    return removed
  }



  removed:=h.elements[0]
  h.elements[0]=h.elements[len(h.elements)-1]

  h.elements=h.elements[:len(h.elements)-1]
  i:=0 
  l:=2*i+1
  r:=2*i+2
  for  2*i+1<len(h.elements){

    if r<len(h.elements) && h.elements[r]<h.elements[l] && h.elements[i]>h.elements[r]{

      h.elements[i],h.elements[r]=h.elements[r],h.elements[i]

      i=r

    }else if h.elements[i]> h.elements[l]{

      h.elements[i],h.elements[l]=h.elements[l],h.elements[i]

      i=l

    }else{
      break
    }


  }

  return removed
}







