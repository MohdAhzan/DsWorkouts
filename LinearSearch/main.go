package main

import (
	"fmt"
	"slices"
)


func main(){
  arr:=[]int {1,3,324,2,34,5,253,2,234,23424,6,56,657,5,5,2125,3,6,4,7,5,2,5,1352,6,6,4,723,1241,242,4,12,4,123,5,1235,7,56,8,635577,3,37,352,36,2}
  slices.Sort(arr)

  fmt.Println(BS(arr,2125))
}


func BS(arr[]int,target int )(int,bool){


  l:=0
  h:= len(arr)-1

  for l<=h {

    m:=(l+h)/2

    if arr[m]==target{
      return arr[m],true
    }
    if target > arr[m]{
      l=m+1
    }else{
      h=m-1
    }

  }

    fmt.Println("no value found")
    return 0,false
}

func linearSearch(arr []int,target int)(bool,int){

  for i,num:=range arr{
    if num == target {
      return true,i
    }
  }
  return false,0
}


