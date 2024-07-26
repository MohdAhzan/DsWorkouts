package queue

import (
	"fmt"
)

type Queue struct{
	
	front *node
	rear *node
}

type node struct  {
	data int
	next *node
}

func (q *Queue)Enqueue(val int){

	newNode:=&node{data: val}

	if q.rear==nil{
		q.rear=newNode
		q.front=newNode
		return
	}

	q.rear.next=newNode
	q.rear=q.rear.next


}

func (q *Queue)Dequeue()int{
	
	if q.front.next==nil{
		fmt.Println("empty queue")
		val:=q.front.data
		q.front=q.front.next
		return val
	}
	val:=q.front.data	
	q.front=q.front.next
	
	return val

}
func (q Queue)PrintQueue(){


	if q.front==nil{
		fmt.Println("empty queue")
		return
	}
	curr:=q.front
	for curr!=nil{
		fmt.Printf("%d  ",curr.data)
		curr=curr.next
	} 
	fmt.Println()
}
