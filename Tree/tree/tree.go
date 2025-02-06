package tree


import "fmt"

type BST struct{
  Root *Node
}
type Node struct{
  lChild *Node
  data int
  rChild *Node
}


func PrintInOrder(curr *Node){
  
  if curr==nil{
    return
  }
  PrintInOrder(curr.lChild)
  fmt.Print(">",curr.data)
  PrintInOrder(curr.rChild)

}

func PrintPostOrder(curr *Node){
  if curr==nil{
    return
  }
  PrintPostOrder(curr.lChild)
  PrintPostOrder(curr.rChild)
  fmt.Print(">",curr.data)
}

func PrintPreOrder(curr *Node){

  if curr==nil{
    return
  }
  fmt.Print("->",curr.data)
  PrintPreOrder(curr.lChild)
  PrintPreOrder(curr.rChild)


}
