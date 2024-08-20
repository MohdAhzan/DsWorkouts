package main

import (
  "ds/Tree/tree"

)


func main() {

  arr:=[]int{12,4,10,32,5,21,22,55}
  // arr:=[]int{4,3,2,5,6,7}
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

  bst.BFS()
}

