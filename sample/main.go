package main

import "fmt"

func main() {
	l:=&Dlist{}

	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Print()
	// l.Delete(1)
	// l.Print()
	// l.Delete(3)
	// l.Print()
	// l.Delete(2)
	// l.Print()
	// l.Delete(4)
	// l.Print()

	l.PrintReverse()
	
	l.Print()
}

type node struct{
	prev *node
	data int
	next *node
}

type Dlist struct{
	head *node
}

func (d Dlist)PrintReverse(){

	if d.head==nil{
		fmt.Println("empty")
		return
	}
	curr:=d.head

	for curr.next!=nil{
		curr=curr.next
	}
	
	tail:=curr

	for tail!=nil{
		

		fmt.Print(tail.data," ")
		tail=tail.prev
	}

fmt.Println()	
}

func (d *Dlist)Add(val int){
	newNode:=&node{data: val}
	if d.head==nil{
		d.head=newNode
		return
	}
	curr:=d.head

	for curr.next!=nil{
		curr=curr.next
	}
	newNode.prev=curr
	curr.next=newNode
}

func (d *Dlist)Delete(target int){
	
	if d.head==nil{
		fmt.Print("empty")	
		return
	}else if d.head.data==target{
		d.head=d.head.next
		if d.head!=nil{
		
			d.head.prev=nil	
		}
		return	
	}

	curr:=d.head
	for curr!=nil&&curr.data!=target{
		curr=curr.next
	}
	if curr == nil{
		return 
	}
	if curr.next == nil{
		curr.prev.next=nil
		return	
	}

	curr.prev.next=curr.next
	curr.next.prev=curr.prev

}

func (d Dlist)Print(){
	if d.head==nil{
		fmt.Println("empty")
		return 
	}	
	curr:=d.head

	for curr!=nil{
		fmt.Print(curr.data," ")
		
		curr=curr.next
	}
	fmt.Println()
}
func checkAnagram(s1,s2 string)bool{

	count:=make(map[rune]int)


	count2:=make(map[rune]int)

	for _,char:=range s1{
		count[char]++
	}

	for _,value:=range s2{
		count2[value]++
	}

	for _,value:=range s2{

		if count[value]!=count2[value] 	{
			return false
		}
	}


	return true
}

func Replace(str, value string, pos int) string {

	char := []byte(str)

	for i:= range char {
		if pos== i+1 {
			char[i]=value[0]
			// char[pos]=
			return string(char)

		}
	}
	return "not found"

}

// type Node struct {
// 	prev *Node
// 	Val  int
// 	next *Node
// }
//
// type Dlink struct {
// 	head *Node
// 	tail *Node
// }
//
// func (l *Dlink) Add(val int) {
//
// 	newNode := &Node{Val: val}
//
// 	if l.head == nil {
// 		l.head = newNode
// 		l.tail = newNode
// 		return
// 	} else {
// 		l.tail.next = newNode
// 		newNode.prev = l.tail
// 		l.tail = newNode
//
// 	}
// }
//
// func (l *Dlink) Print() {
//
// 	curr := l.head
//
// 	for curr != nil {
// 		fmt.Println(curr.Val)
// 		curr = curr.next
//
// 	}
// 	fmt.Println()
// }
