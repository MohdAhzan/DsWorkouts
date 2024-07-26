package main

import (
	"fmt"
)

func main() {

	arr:=[]int{5,2,5,3,6,2,26,36,2,2,3,5,53,2,33,5,435,432,53,42,53,5,52,34,23,445,3}

	MergeSort(arr)

	fmt.Println(arr)


	// s:=StackInit()
	// s.Push(4)
	// s.Push(6)
	// s.Push(9)
	// s.Push(3)
	// s.Push(1)
	//
	// s.PrintStack()
	//
	// s.ReverseStack()
	// s.PrintStack()
	// curr:=s.Top
	// q:=QueueInit()	
	// for curr!=nil{
	// 	q.Enqueue(curr.data)	
	// 	// fmt.Print(curr.data)
	// 	curr=curr.next
	// }
	// q.PrintQueue()
	
}

func QuickSort(arr []int){

	partion(arr,0,len(arr)-1)
	return
}

func partion(arr[]int,start,end int){
	
	if start > end{
		return
	}

	pivot:=0
	left:=start+1
	right:=end
	
	for left <= right{
	
		if arr[left]>arr[pivot]&&arr[right]<arr[pivot]{

			arr[left],arr[right]=arr[right],arr[left]
			left++
			right--
		}

		if arr[left]<arr[pivot]{
			left++
		}
		if arr[right]>arr[pivot]{
			right++
		}


	}

	arr[pivot],arr[right]=arr[right],arr[pivot]

	partion(arr,start,right-1)
	partion(arr,right+1,end)
}

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
			
			if arr[j]<arr[smIndx]{

				smIndx=j
			}
		}
		arr[i],arr[smIndx]=arr[smIndx],arr[i]

	}
}

func MergeSort(arr []int){
	
	m:=len(arr)/2

	if len(arr)<=1{
		return
	}

	leftArr:=make([]int,m)
	rightArr:=make([]int,len(arr)-m)

	copy(leftArr,arr[:m])
	copy(rightArr,arr[m:])

	MergeSort(leftArr)
	MergeSort(rightArr)

	i,j,k:=0,0,0

	for i<len(leftArr)&&j<len(rightArr){
	
		if leftArr[i]<rightArr[j]{
			arr[k]=leftArr[i]
			i++
		
		}else{
			arr[k]=rightArr[j]

			j++
		}
		k++

	}

	for i<len(leftArr){

		arr[k]=leftArr[i]
		i++
		k++
	}
	for j<len(rightArr){
		arr[k]=rightArr[j]
		j++
		k++
	}

}

func (s Stack)PrintStack(){

	if s.Top==nil{
		fmt.Println("empty stack")
		return
	}
	curr:=s.Top

	for curr!=nil{
		fmt.Printf("%d\t", curr.data)
		curr=curr.next
	}
	fmt.Println()
}



type Node struct{

	data int
	next *Node
}
type Stack struct{

	Top *Node
}

func StackInit()*Stack{
	
	return &Stack{}

}

func (s *Stack)ReverseStack(){
	var prev,next *Node
	curr:=s.Top
	for curr!=nil{

		next=curr.next
		curr.next=prev
		prev=curr
		curr=next
	}

	s.Top=prev	
}

func (s *Stack)Push(val int){
	
	newNode:=&Node{data: val}

	if s.Top==nil{
		s.Top=newNode
		return
	}
	newNode.next=s.Top	
	s.Top=newNode
}

func (s *Stack)Pop(){

	if s.Top==nil{
		fmt.Println("stack Underflow")
		return 
	}

	s.Top=s.Top.next
}

type Queue struct{
	
	front *Node
	rear *Node

}
func (q *Queue)Enqueue(val int){

	newNode:=&Node{data: val}

	if q.front==nil{
		q.front=newNode
		q.rear=newNode
		return
	}
	q.rear.next=newNode
	q.rear=newNode
}

func (q Queue)PrintQueue(){

	if q.front==nil{
		return
	}
	curr:=q.front

	for curr!=nil{
		fmt.Printf("%v\t", curr.data)
		curr=curr.next
	}
	fmt.Println()
}
func QueueInit()*Queue{
	return&Queue{}
}
