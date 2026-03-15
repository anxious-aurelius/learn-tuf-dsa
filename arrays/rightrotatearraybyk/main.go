package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(rotateright(arr, 2))
	//fmt.Print(rightrotatebyone(arr))
}

func rotateright(arr []int, k int) []int {
	k = k % len(arr)
	reversearray(arr)
	reversearray(arr[:k])
	reversearray(arr[k:])
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

//func rightrotatebyone(arr []int) []int {
//	temp := arr[len(arr)-1]
//	for i := len(arr) - 1; i > 0; i-- {
//		arr[i] = arr[i-1]
//	}
//	arr[0] = temp
//	return arr
//}
