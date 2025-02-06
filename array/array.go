package main

import (
	"fmt"
)


func FindNumberSum(arr []int, target int) (a, b int) {
  set := make(map[int]struct{})

  for i := 0; i < len(arr); i++ {
    match := target - arr[i]
    fmt.Println(match,"matcH")
    _,ok:=set[match]
    if ok {
      return arr[i], match
    } else {

      set[i] = struct{}{} 
    }

  }

  fmt.Println(set,set[1])
  return 0, 0
}


func SendToEnd(a[]int,target int){
  arr:=a
  i:=0
  j:=len(arr)-1

  for i<len(arr)-1{
    fmt.Println(i,j)	
    if i>j{
      break
    }
    if arr[j]==target{
      j--	
    }
    if arr[i] == target {		
      arr[i],arr[j]= arr[j],arr[i]
      i++
      j--
    }
    i++		
  }
  fmt.Print(arr)	
}



