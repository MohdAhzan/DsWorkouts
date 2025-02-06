package main

import (
	"fmt"
)

func main() {

	arr := []int{1, 6, 3, 7, 2, 44, 23, 1, 5, 16, 72, 22, 32}

	fmt.Println(arr)

	Quicksort(arr)

	fmt.Print(arr)
}

func Quicksort(arr []int){


	partion(arr,0,len(arr)-1)
	return
}
func partion(arr[]int , start,end int){

	if start > end {
		return
	}

	pivot:=start

	left:=start+1 
	right:=end

	for left<=right{

		if arr[left]>arr[pivot]&&arr[right]<arr[pivot]{
		
			arr[left],arr[right]=arr[right],arr[left]

			left++
			right--
		}

		if arr[left]<=arr[pivot]{

			left++
		}

		if arr[right]>=arr[pivot]{
			right--
		}


	}

	arr[pivot],arr[right]=arr[right],arr[pivot]

	partion(arr,start,right-1)
	partion(arr,right+1,end)


}


func MergeSort(arr []int) {

	m := len(arr) / 2

	if len(arr) <= 1 {
		return
	}

	leftArr := make([]int, m)
	rightArr := make([]int, len(arr)-m)

	copy(leftArr, arr[:m])
	copy(rightArr, arr[m:])
	

	MergeSort(leftArr)
	MergeSort(rightArr)

	i, j, k := 0, 0, 0

	for i < len(leftArr) && j < len(rightArr) {

		if leftArr[i] < rightArr[j] {

			arr[k] = leftArr[i]
			i++
		} else {
			arr[k] = rightArr[j]
			j++
		}
		k++
	}

	for i < len(leftArr) {
		arr[k] = leftArr[i]
		i++
		k++
	}
	for j < len(rightArr) {
		arr[k] = rightArr[j]
		j++
		k++
	}
}

func InsertionSort(arr []int) {

	for i := 1; i < len(arr); i++ {

		j := i - 1

		for j >= 0 && arr[j+1] < arr[j] {

			temp := arr[j+1]
			arr[j+1] = arr[j]
			arr[j] = temp

			j = j - 1
		}

	}

}

func selectionSOrt(arr []int) {

	for i := 0; i < len(arr); i++ {

		smallindex := i
		for k := i + 1; k < len(arr); k++ {
			if arr[smallindex] > arr[k] {
				smallindex = k
			}
		}

		arr[i], arr[smallindex] = arr[smallindex], arr[i]

	}
}

func BubbleSort(arr []int) {

	for i := 0; i < len(arr)-1; i++ {

		for j := 0; j < len(arr)-i-1; j++ {

			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]

			}
		}
	}
}
