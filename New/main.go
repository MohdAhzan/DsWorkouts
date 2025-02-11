package main

import "fmt"

func main() {

  //   b:=&list{}
  //
  // b.Insert(4)
  // b.Insert(13)
  // b.Insert(61)
  // b.Insert(9)
  //
  // b.Print()
  //
  // b.InsertAfterPOS(4,1)
  // b.Print()
  //
  // b.InsertAfterPOS(1,888)
  // b.Print()
  //
  // b.InsertAfterPOS(9,0)
  // b.Print()



  grid:=[][]int{
    {0,0,1,1,0},
    {0,0,0,1,0},
    {0,0,0,0,0},
    {1,1,1,1,0},
  }
  visit:=make(map[[2]int]bool)


  routes :=DFS(grid,0,0,visit)

  fmt.Println("routes to travel in the grid are ",routes)

}

func DFS(grid [][]int, r,c int, visit map[[2]int]bool)int{

  ROWS:=len(grid)
  COLS:=len(grid[0])

  if r<0 || c < 0 ||  visit[[2]int{r,c}] || r == ROWS || c == COLS || grid[r][c]==1 {

    return 0
  }
  if r==ROWS-1 && c == COLS-1{

    return 1
  }
  visit[[2]int{r,c}]=true

  count:=0

  count+= DFS(grid,r+1,c,visit)
  count+= DFS(grid,r-1,c,visit)
  count+= DFS(grid,r,c+1,visit)
  count+= DFS(grid,r,c-1,visit)

  visit[[2]int{r,c}]=false

  return count

}

type list struct{

  head *node

}

type node struct{

  prev *node
  Val int
  next *node

}


func (b *list)Print(){
  
  curr:=b.head

  for curr!=nil{
  
    fmt.Print(" < ",curr.Val," > " )
    curr=curr.next
  }
  fmt.Println()

}

func (b *list)InsertAfterPOS(pos,data int){
  
  newNode:=&node{Val: data}
  
  if b.head==nil{
    fmt.Println("empty") 
    return
  }else if b.head.Val==pos{
  
    newNode.prev =b.head
    newNode.next=b.head.next
    b.head.next=newNode
    return
  }

  curr:=b.head

  for curr.next!=nil && curr.next.Val != pos{
  
    curr=curr.next
    
  }

  if curr.next==nil {
    
    fmt.Println("no posValue found in list")
    return
  }else {
      newNode.next=curr.next.next
      newNode.prev=curr.next
      curr.next.next=newNode
  }
  
    
}
func (b *list)Insert(data int){

  newNode:=&node{Val: data}

  if b.head==nil{
    b.head=newNode
    return 
  }

  curr:=b.head 

  for curr.next!=nil {
        curr=curr.next 
    }

  newNode.prev=curr
  curr.next=newNode

}
