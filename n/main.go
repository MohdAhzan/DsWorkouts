package main

func main(){
}

type Trie struct{

  Root *node
}

type node struct{
  
  child map[rune]*node
  isEnd bool
}





