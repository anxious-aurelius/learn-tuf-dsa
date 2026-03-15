package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(leftrotatebyk(arr1, 2))
}

func leftrotatebyk(arr []int, k int) []int {
	k = k % len(arr)
	reversearray(arr[:k])
	reversearray(arr[k:])
	reversearray(arr)
	return arr
}

func reversearray(arr []int) {
	left := 0
	right := len(arr) - 1
	for left < right {
		temp := arr[left]
		arr[left] = arr[right]
		arr[right] = temp
		left++
		right--
	}
	return
}

//Higher Time Complexity
//func leftrotatebyk(arr []int, k int) []int {
//	k = k % len(arr)
//	if len(arr) == 0 {
//		return []int{}
//	} else if len(arr) == 1 {
//		return arr
//	}
//	for i := 0; i < k; i++ {
//		temp := arr[0]
//		for j := 1; j < len(arr); j++ {
//			arr[j-1] = arr[j]
//		}
//		arr[len(arr)-1] = temp
//	}
//	return arr
//}
