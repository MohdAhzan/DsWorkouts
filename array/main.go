package main


func main(){
arr:=[][]int{}
}

func result(arr[]int, target int)(int,int){
	for i:=0;i<len(arr)-2;i++{
		for j:=i+1;j<len(arr)-1;j++{
			if arr[i]+arr[j] == target{
				return arr[i],arr[j]
			}
		}
	}
	return 0,0
}
