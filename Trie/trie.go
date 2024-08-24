package main

import (
  "fmt"
)

func main(){
  
  t:=NewTrie() 
  t.Insert("ahzan")
  t.Insert("ahz")
  t.Insert("ah")
  t.Insert("hmm")
  t.Insert("hm")
  
  // t.PrintAll()

  t.PrintwordswithPrefix("a")

  fmt.Println(len(t.root.child)) 
  
}

type Trie struct {
  
  root *node

}

type node struct{
  
  child map[rune]*node
  isEnd bool

}

func NewTrie()*Trie{
  return  &Trie{root: &node{child: make(map[rune]*node)}}
}

func (t *Trie)Insert(word string){
  
curr:=  t.root

  for _,char:=range word {
    
    if _,exists:= curr.child[char]; !exists{

      curr.child[char]=&node{child: make(map[rune]*node)}
    }
    curr=curr.child[char]

  }
  
  curr.isEnd=true
    
}

func (t *Trie)PrintAll(){
    
  t.collectWord(t.root,"")

}

func (t *Trie)collectWord(node *node,prefix string){
  
    if node.isEnd{
      fmt.Println(prefix)
  }
    for ch,childNode:=range node.child{
  
        t.collectWord(childNode,prefix+string(ch))
      
  }
}

func (t *Trie)PrintwordswithPrefix(prefix string){
  
    curr:=t.root
  
    for _,ch:=range prefix{
    
    if _,ok:=curr.child[ch];!ok{
      
      fmt.Println("No words found with prefix :",prefix)
        return

    }
    curr=curr.child[ch]

  } 
    
  t.collectWord(curr,prefix)
}


