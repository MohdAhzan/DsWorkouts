package stack

import "fmt"


type Stack struct{
	
	Top *node
	Size int

}

type node struct{
	data int
	next *node 
}

func StackInitialize()*Stack{

	return &Stack{}
}

func (s *Stack)PrintStack(){
	curr:=s.Top

	for curr!=nil{
		fmt.Printf("%d\n",curr.data)
		curr=curr.next
	}
	fmt.Println()
}

func (s *Stack)Push(value int){

	newNode:=&node{data:value}
	// if s.Size<=0{
	// 	fmt.Println("stack overflow")	
	// 	return
	// }
	if s.Top==nil{
		s.Top=newNode

		return
	}else{
		newNode.next=s.Top
		s.Top=newNode
		return
	}


}

func (s *Stack)Pop(){

	if s.Top==nil{
		fmt.Println("stack underflow")
		return
	}
		s.Top=s.Top.next
		

}



