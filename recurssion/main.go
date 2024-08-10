package main

import "fmt"

func main() {

  arr := []int{10, 20, 30, 40, 50, 60, 70}

  index := BinarySearch(arr, 90, 0, len(arr)-1)
  fmt.Println(index)

  str := "ahzan"
  str2 := "malayalam"

  fmt.Println(Reverse(str, len(str)))
  fmt.Println(Reverse(str2, len(str2)))
  // fmt.Println(PalindromeCheck(str,0,len(str)-1))

  // fmt.Println(PalindromeCheck(str2,0,len(str2)-1))
  // fmt.Println(Reverse(str,len(str)))
  // sum := sumREc(arr, 0)
  // fmt.Println("sum", sum)
  //
  // fac := Factorial(5)
  //
  // fmt.Println("factorial", fac)
}

func Reverse(s string, n int) string {

  if n == 0 {
    return ""
  }

  return string(s[n-1]) + Reverse(s, n-1)

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

func PalindromeCheck(s string, left, right int) bool {

  if left > right {
    return true
  }
  if s[left] != s[right] {
    return false
  }
  return PalindromeCheck(s, left+1, right-1)

}

func BinarySearch(arr []int, target, low, high int) int {

  if low > high {
    fmt.Println("no value found")
    return -1
  }

  mid := (low + high) / 2

  if arr[mid] == target {
    return mid
  }
  if arr[mid] > target {
    return BinarySearch(arr, target, low, mid-1)
}else {
      return BinarySearch(arr, target, mid+1, high)
    }


 }
