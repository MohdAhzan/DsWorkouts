package main

import "fmt"

func main() {

	l := &list{}

	l.Add(9)
	l.Add(2)
	l.Add(4)
	l.Add(9)
	l.Add(3)
	l.print()

	l.Delete(9)
	l.Delete(3)
	l.Delete(9)
	l.print()

	l.Delete(4)
	l.print()

	l.Delete(2)
	l.print()

	l.Delete(2)
	l.print()

}

type list struct {
	head *node
}
type node struct {
	prev *node
	val  int
	next *node
}

func (l *list)Delete(target int){

	if l.head==nil{
		fmt.Println("deleting value from an empty list is not possible you dumb asss motherfuckerrr....!!!!!! ")
		return
	}else if l.head.val==target{


		if l.head.next!=nil{
		l.head=l.head.next
		l.head.prev=nil
		return
		}
		l.head=nil
		return

		}
	

	curr:=l.head.next

	for curr!=nil&&curr.val!=target{
		curr=curr.next
	}

	if curr==nil{
		return
	}
	if curr.next==nil&&curr.val==target{
	
		curr.prev.next=nil
		curr=nil
		return
	}

	curr.prev.next=curr.next
	curr.next.prev=curr.prev
	
}


func (l *list) Add(value int) {
	newList := &node{val: value}

	if l.head == nil {
		l.head = newList
		return
	}
	//  else {
	// 	newList.prev = l.head
	// 	l.head.next = newList
	// 	return
	// }

	curr := l.head

	for curr.next != nil {
		curr = curr.next

	}

	if curr.next == nil {
		newList.prev = curr
		curr.next = newList
	}

}

func (l list) print() {

	if l.head == nil {
		
		return
	}

	c := l.head

	for c != nil {

		fmt.Println(c.val)

		c = c.next
	}
	fmt.Println()
}
