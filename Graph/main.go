package main

import (
  "fmt"
)

func main(){

  grid:=[][]int{
    {0,0,0,0},
    {1,1,0,0},
    {0,0,0,1},
    {0,1,0,0},
  }
  visit:=make(map[[2]int]bool)

  routes:= GraphDFS(grid,0,0,visit)
  fmt.Println("possible routes in this graph is ",routes)

}

func GraphDFS(grid [][]int ,r,c int,visit map[[2]int]bool )int{
  
  
  ROWS,COLS:=len(grid),len(grid[0])

  if r<0 ||c < 0 || r==ROWS ||c == COLS || visit[[2]int{r,c}]|| grid[r][c]==1 {
    return 0
  }
  if r== ROWS-1 && c == COLS-1 {
    return 1
  }

  visit[[2]int{r,c}]=true
  
  count:=0
  count =GraphDFS(grid,r+1,c,visit)+count
  count =GraphDFS(grid,r-1,c,visit)+count
  count =GraphDFS(grid,r,c-1,visit)+count
 count =GraphDFS(grid,r,c+1,visit)+count

  visit[[2]int{r,c}]=false
  
  return count
}
