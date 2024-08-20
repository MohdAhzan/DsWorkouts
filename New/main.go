package main

import "fmt"

func main() {

  //
  // grid:=[][]int{
  //   {0,0,0,0},
  //   {0,1,0,0},
  //   {0,0,0,1},
  //   {0,1,0,0},
  // }
  //


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

type BST struct{

  Root *node

}

type node struct{

  Left *node
  Val int
  Right *node

}

func (b *BST)Insert(data int){

  newNode:=&node{Val: data}

  if b.Root==nil{

    b.Root=newNode
    return 

  }

  curr:=b.Root 

  for curr!=nil {

    if data < curr.Val{


    }

  }



}
