package tree


func (r *BST)InsertTree(val int){
    
  newNode:=&Node{data: val}
  if r.Root==nil{
  
    r.Root=newNode

  }else{
    insertNodes(r.Root,newNode)
  }

}

func insertNodes(curr,newNode *Node){
    
  if curr.data>newNode.data{
  
    if curr.lChild==nil{
      curr.lChild=newNode
    }else{
        
      insertNodes(curr.lChild,newNode)

    }
      
  }else if curr.data <newNode.data{

        if curr.rChild==nil{
      curr.rChild=newNode
    }else{
      insertNodes(curr.rChild,newNode)
    }

  }
}


func (r *BST)InsertItrerative(value int){
  
  newNode:=&Node{data: value}
  curr:=r.Root
  if curr==nil{
    r.Root=newNode
    return
  }
  
  for curr!=nil{
  
      if value > curr.data{
 
      if curr.rChild==nil{
        curr.rChild=newNode
        return
      }else{
        curr=curr.rChild
      }
    }else if value < curr.data{
  
        if curr.lChild==nil{
        curr.lChild=newNode
        return
      }else{
        curr=curr.lChild
      }
    }

  }


}


