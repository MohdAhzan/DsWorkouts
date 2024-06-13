package main

import (
	"fmt"
	"slices"
)

func main(){

	arr:=[]int {1,3,324,2,34,5,253,2,234,23424,6,56,657,5,5,2125,3,6,4,7,5,2,5,1352,6,6,4,723,1241,242,4,12,4,123,5,1235,7,56,8,635577,3,37,352,36,2}

	slices.Sort(arr)
	fmt.Println("sorted array",arr)
	fmt.Println(binarySearch(arr,1352))
}
func  binarySearch(arr []int,target int) (bool,int){


	if len(arr) == 0 {
		fmt.Println("Empty array")
		return false,0
	}
	low:=0
	high:=len(arr)-1

	for low <= high {

		median:= low+(high-low)/2

		if arr[median]== target {
			return true,median
		}
		if arr[median]<target{
			low=median+1
		}else{
			high=median-1
		}


	}

	return false,0
}
