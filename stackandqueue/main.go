package main

import (
	"ds/stackandqueue/queue"
	"ds/stackandqueue/stack"
)



func main(){

  
	q:=&queue.Queue{}
	q.Enqueue(7)
	q.Enqueue(4)
	q.Enqueue(0)
	q.Enqueue(2)
	q.PrintQueue()
	
	s:=stack.StackInitialize()	
	s.Push(q.Dequeue())
	
	s.PrintStack()

}
var arr []int

func PushArr(value int, arr[]int  ){

	n:=len(arr)
	if n==0{
		arr[0]=value
		return
	}

	arr[n-1]=0
}

func PopArr(arr []int){
	
	arr=arr[:len(arr)-1]

}

