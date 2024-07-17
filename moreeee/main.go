package main

import (
	"fmt"
	"slices"
)

func main() {
	s := "ahzan"
	fmt.Println(reverse(s))

	n := 5

	fmt.Println(factorial(n))

	ar := []int{1, 34, 234, 2, 4, 234, 2, 4, 25, 23, 5, 4354, 6, 34, 63, 46, 54, 6, 3, 424, 23, 4, 25, 3}
	slices.Sort(ar)
	fmt.Println(ar)
	fmt.Println()
	binarysearch(ar, 0, len(ar)-1, 25)

}

func reverse(s string) string {

	if len(s) == 0 {
		return ""
	}

	return string(s[len(s)-1]) + reverse(s[:len(s)-1])

}

func factorial(n int) int {

	if n == 1 {
		return n
	}
	return n * factorial(n-1)

}

func binarysearch(arr []int, left, right, target int) int {

	mid := (left + right)/2

	if left > right {
		fmt.Println("no value found")
		return -1
	}

	if arr[mid] == target {
		fmt.Println("value at index", mid)
		return mid
	}
	if target < arr[mid] {
		return binarysearch(arr, left, mid-1, target)
	}

	return binarysearch(arr, mid+1, right, target)

}
