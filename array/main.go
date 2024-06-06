package main


func main(){
	//a:=[]int{1,3,5,6,7,2,4}
	//t:=10
	//x,y:=result(a,t)
	arrA:=[]int{3,4,6,23,2,5,6,9,4,23,8,4,2,3,4,6,8,0,5,3,2,9,4}
	//x,y:=FindNumberSum(a,t)
	SendToEnd(arrA,4)
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
