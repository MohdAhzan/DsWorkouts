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

func RemoveTest (root *Node,target int)*Node{


  if root==nil{
    return nil
  }

  if target<root.data{
  
      root.lChild=RemoveTest(root.lChild,target)
  }else if target > root.data{
    root.rChild=RemoveTest(root.rChild,target)
  }else{
  
    if root.lChild==nil{
      return root.rChild
    }else if root.rChild==nil{
      return root.lChild
    }else{
  
      min:=MinNode(root.rChild) 
      root.data=min.data
      root.rChild=RemoveTest(root.rChild,min.data)


    }

  }

  return root
}

func MinNode(root *Node)*Node{
  curr:=root 

  for curr!=nil&&curr.lChild!=nil{
     curr=curr.lChild 

  }
  return curr
}
