package tree
import "fmt"
func BooleanSearchTree(curr *Node,val int)bool{

  if curr==nil{
    return false
  }
  if curr.data>val{
    return BooleanSearchTree(curr.lChild,val)
  }else if curr.data < val{
    return BooleanSearchTree(curr.rChild,val)
  }

  return true


}

func (r *BST)SearchIterative(val int)bool{

  if r.Root ==nil{
    fmt.Println("empty BST")
    return false
  }  
  curr:=r.Root

  for curr!=nil{
  
        if val < curr.data{
          curr=curr.lChild
    }else if val > curr.data{
      curr=curr.rChild
    }else{
      return true
    }
  } 
  
    return false
}

