package main

import (
  "ds/Tree/tree"
  "fmt"

)


func main() {



  // arr:=[]int{12,4,10,32,5,21,22,55}
  arr:=[]int{4,3,2,5,6,7}
  bst:=&tree.BST{}
  for i := range arr {

    bst.InsertItrerative(arr[i])
    //   12
    //  /  \
    // 4    32
    //  \
    //   10
    //  /
    // 5
    //  \
    //   8
  } 

  tree.PrintInOrder(bst.Root)  

  if  bst.SearchIterative(5){
    fmt.Println("found")
  }else{
    fmt.Println("not  found")
  }
  r:= tree.RemoveTest(bst.Root,4)

  tree.PrintInOrder(bst.Root)
  fmt.Println("r",r)
}

