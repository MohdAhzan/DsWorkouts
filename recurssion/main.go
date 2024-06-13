package main

import "fmt"

func main() {

	arr := []int{10, 20, 30, 40, 50, 60, 70}

	sum := sumREc(arr, 0)
	fmt.Println("sum", sum)

	fac := Factorial(5)

	fmt.Println("factorial", fac)
}

func sumREc(arr []int, i int) int {

	if i == len(arr) {
		return 0
	}

	return arr[i] + sumREc(arr, i+1)

}

func Factorial(num int) int {

	if num <= 1 {
		return num
	}

	return num * Factorial(num-1)

}
