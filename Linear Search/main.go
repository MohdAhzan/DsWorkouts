package main

import "fmt"


func main(){
	arr:=[]int {1,3,324,2,34,5,253,2,234,23424,6,56,657,5,5,2125,3,6,4,7,5,2,5,1352,6,6,4,723,1241,242,4,12,4,123,5,1235,7,56,8,635577,3,37,352,36,2}

	fmt.Println(linearSearch(arr,1352))
	fmt.Println(arr[23])
}

func linearSearch(arr []int,target int)(bool,int){
	
	for i,num:=range arr{
		if num == target {
			return true,i
		}
	}
return false,0
}


