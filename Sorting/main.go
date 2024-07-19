package main

import "fmt"

func main(){

	arr:=[]int{1,6,3,7,2,44,23,1,5,16,72,22,325,3634,634,323,1,421,4253,63,34,23,41214,2,3634,434,3}
	
	
	BubbleSort(arr)
	fmt.Print(arr)

}


func InsertionSOrt(arr []int){

	for i:=0;i<len(arr);i++{

		j:=i-1
	
		for j>=0&&arr[j+1] < arr[j]{

			temp:=arr[j+1]
			arr[j+1]=arr[j]
			arr[j]=temp
			j=j-1
	
		}
	}
}	


func BubbleSort(arr []int){


	for i:=0;i<len(arr)-1;i++{

		for j:=0;j<len(arr)-i-1;j++{

			if arr[j]>arr[j+1]{
			arr[j],arr[j+1]=arr[j+1],arr[j]

			}
		}
	}
}


