package main

import (
	"container/list"
	"fmt"
)

func main() {
	b := &BinaryTree{}

	// b.root=&Node{data: 14}
	// b.root.left=&Node{data: 13}
	// b.root.right=&Node{data: 20}
	// b.root.right.left=&Node{data: 22}
	// b.root.right.right=&Node{data: 2}
	//

	arr := []int{12, 2, 34, 4, 56, 7, 78, 3}
	//        12
	//    /      \
	//    2      34
	//    \       \
	//     4       56
	//   /  \        \
	//  3     7        78
	for i := range arr {
		b.Insert(arr[i])

	}
  BFSNew(b.root)
  // printInorder(b.root)
	fmt.Println()

	fmt.Println(CheckBST(b.root))
}

func CheckBST(curr *Node) bool {
	if curr == nil {
		return true
	}

	if curr.left != nil && MaxValue(curr.left) >= curr.data {
		return false
	}

	if curr.right != nil && MinValue(curr.right) <= curr.data {
		return false
	}

	return CheckBST(curr.left) && CheckBST(curr.right)

}
func MaxValue(curr *Node) int {

	for curr.right != nil {
		curr = curr.right
	}
	return curr.data

}
func MinValue(curr *Node) int {

	for curr.left != nil {
		curr = curr.left
	}
	return curr.data

}
func (b *BinaryTree) getvalue(target int) int {

	return helperSearch(b.root, target)

}

func helperSearch(curr *Node, target int) int {

	if curr == nil {
		fmt.Println("value not founnd")
		return -1
	}
	if target < curr.data {

		return helperSearch(curr.left, target)
	} else if target > curr.data {

		return helperSearch(curr.right, target)
	}

	fmt.Println("value found ")
	return curr.data

}

func printInorder(curr *Node) {

	if curr == nil {
		return
	}

	printInorder(curr.left)
	fmt.Print(curr.data, "---")
	printInorder(curr.right)

}

type BinaryTree struct {
	root *Node
}

type Node struct {
	left  *Node
	data  int
	right *Node
}

func (b *BinaryTree) Insert(val int) {

	newNode := &Node{data: val}

	if b.root == nil {
		b.root = newNode
		return
	}
	curr := b.root

  for curr!=nil{
  
      if val < curr.data{

      if curr.left==nil{
        curr.left=newNode
        return
      }else{
        curr= curr.left
      }
    }else if val > curr.data{

      if curr.right==nil{
        curr.right=newNode
        return
      }else{
        curr= curr.right
      }
    }

  }
}

func (b *BinaryTree) InsertTree(val int) {

	newNode := &Node{data: val}
	if b.root == nil {
		b.root = newNode

	} else {

		InsertNodes(b.root, newNode)
	}

}

func InsertNodes(curr, newNode *Node) {

	if newNode.data < curr.data {
		if curr.left == nil {
			curr.left = newNode
		} else {
			InsertNodes(curr.left, newNode)
		}

	} else if newNode.data > curr.data {

		if curr.right == nil {
			curr.right = newNode
		} else {
			InsertNodes(curr.right, newNode)
		}

	}

}

func Remove(root *Node, target int) *Node {

	if root == nil {
		return nil

	}
	if target < root.data {

		root.left = Remove(root.left, target)
	} else if target > root.data {
		root.right = Remove(root.right, target)
	} else {

		if root.right == nil {
			return root.left
		} else if root.left == nil {

			return root.right
		} else {

			minNode := Min(root.right)
			root.data = minNode.data
			root.right = Remove(minNode.right, minNode.data)

		}

	}
	return root
}

func Min(root *Node) *Node {

	curr := root

	for curr != nil && curr.left != nil {

		curr = curr.left
	}
	return curr
}

func BFS(root *Node) {

	if root == nil {
		return
	}

	queue := list.New()
	queue.PushBack(root)
	// level:=0
	for queue.Len() > 0 {

		// fmt.Print("level ",level,": ")
		// levelLength:=queue.Len()
		// for i:=0;i<levelLength;i++{

		curr := queue.Front().Value.(*Node)
		queue.Remove(queue.Front())
		fmt.Print(curr.data, ">")

		if curr.left != nil {
			queue.PushBack(curr.left)
		}
		if curr.right != nil {
			queue.PushBack(curr.right)
		}

		// }
		// level++
		// fmt.Println()

	}

}


func BFSNew(root *Node) {

	if root == nil {
		return
	}

  
    queue:=list.New()
    queue.PushBack(root)



    for queue.Len()>0{
        levelLength:=queue.Len()

        for i:=0;i<levelLength;i++{

      curr:=queue.Front().Value.(*Node)
      queue.Remove(queue.Front()) 
      fmt.Print(curr.data," ")

      if  curr.left!=nil{
        queue.PushBack(curr.left)
      }
      if curr.right!=nil{
        queue.PushBack(curr.right)
      }

    }
    fmt.Println()
        
  }

}

