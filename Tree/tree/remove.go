package tree

func MinValueNode(root *Node)*Node{

  curr:=root 

  for curr!=nil && curr.lChild!=nil{

    curr=curr.lChild
  }

  return curr 
}

func Remove(root *Node,val int)*Node{


  if root==nil{

    return nil
  }
  if val > root.data{
    root.rChild=Remove(root.rChild,val)
  }else if val < root.data{
    root.lChild=Remove(root.lChild,val)
  }else{
    if root.lChild ==nil{
      return root.rChild
    }else if root.rChild==nil{
      return root.lChild
    }else{
      minNode:=MinValueNode(root.rChild)
      root.data=minNode.data
      root.rChild = Remove(root.rChild,minNode.data)
    }

  }
  return root
}

func RemoveTest (curr *Node,target int)*Node{
  
    if curr==nil{
    return nil
  }

  if target > curr.data{
  
      curr.rChild=RemoveTest(curr.rChild,target)
       
        
  }else if target < curr.data{
    
    curr.lChild=RemoveTest(curr.lChild,target)

  }else{
      
      if curr.lChild==nil{
        return curr.rChild 
    }else if curr.rChild==nil{
      return curr.lChild
    }else{
    
        min:=MinNode(curr.rChild)
        curr.data=min.data
        curr.rChild=RemoveTest(curr.rChild,min.data)
            
    }

      
  }
  
  return curr
    
}

func MinNode(root *Node)*Node{
  
    for root !=nil && root.lChild!=nil{

    root=root.lChild
  }
  return root
}
