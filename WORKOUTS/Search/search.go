package main

import (
	"fmt"
)

func main() {

  st:="Ahzan"

  fmt.Println(st[1:])

  
}




func Search(arr[]int,t int)(bool,int){

  for _,n := range arr {
    if n==t{
      return true,n
    }
  }
  return false,0
}

func BinarySearchRecursion(arr[]int, l,h,target int)(bool,int){

   if l>h {
    return false ,0
  }

    mid:=l+(h-l)/2
  
  if arr[mid]==target{
    return true,arr[mid]
  }
  if arr[mid]>target{
     return BinarySearchRecursion(arr,l,mid-1,target)
  }else {
     return BinarySearchRecursion(arr,mid+1,h,target)
  }
}

func BinarySearch(arr[]int,t int)(bool,int){

  i:=0
  j:=len(arr)-1

  for i<=j{

    mid:=(i+j)/2
    fmt.Println(mid, i, j)
    if arr[mid]==t{
      return true,arr[mid]
    }else if arr[mid]<t{
      i=mid+1 
    }else{
      j=mid-1
    }
  }
  return false,0
}

func Reverse(s string)string{
 
    STR:=[]byte(s)
    fmt.Println("str in bytes",STR)
    i:=0
    j:=len(STR)-1

  for i<j{

    STR[i],STR[j]=STR[j],STR[i]
    i++
    j--
  }
  return string(STR)
}

func RecursionReverse(s string)string{
  if len(s)==0{
    return s
  }
  return RecursionReverse(s[1:])+string(s[0])
}


