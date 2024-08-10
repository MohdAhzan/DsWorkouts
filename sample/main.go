package main

import (
	"fmt"
)

func main() {

	arr:=[]int{12,34,22,11,1,6,3,21,99,0}
	
	fmt.Println( arr)
	SelectionSort(arr)	
	fmt.Println(arr)
}


// func BubbleSort(arr []int){
//
// 	n:=len(arr)
//
// 	for i:=0;i<n-1;i++{
// 		for j:=0;j<n-i-1;j++{
// 			if arr[j]>arr[j+1]{
//
// 				arr[j],arr[j+1]=arr[j+1],arr[j]
//       }
// 			}
//
// 	}
//
// }
//
func InsertionSort(arr []int){

	n:=len(arr)
	for i:=1;i<n;i++{

		j:=i-1
		for j>=0&&arr[j+1]<arr[j]{

			arr[j+1],arr[j]=arr[j],arr[j+1]
			j--	
		}
	}
}

func SelectionSort(arr []int){
	
	n:=len(arr)
	for i:=0;i<n-1;i++{

		smIndx:=i

		for j:=i+1;j<n;j++{

			if arr[j]< arr[smIndx] {
				smIndx=j
			}
		}

		arr[i],arr[smIndx]=arr[smIndx],arr[i]
	}
}


