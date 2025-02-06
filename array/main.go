package main

import "fmt"


func main(){

  arr:=[]int{1,5,3,2,5}
    // SendToEnd(arr,1)

  SendtoLastwithoutOrderChange(arr)

    fmt.Println("arr",arr)

}

func result(arr[]int, target int)(int,int){
	for i:=0;i<len(arr)-2;i++{
		for j:=i+1;j<len(arr)-1;j++{
			if arr[i]+arr[j] == target{
				return arr[i],arr[j]
			}
		}
	}
	return 0,0
}

func SendtoLastwithoutOrderChange(arr []int){
  j:=0
  for i:=1;i<len(arr);i++{

     arr[i],arr[j]=arr[j],arr[i]

    j++

  }
    

}
