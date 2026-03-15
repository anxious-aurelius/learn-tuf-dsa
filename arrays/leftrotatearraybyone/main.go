package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(leftrotatebyone(arr1))
}

func leftrotatebyone(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	} else if len(arr) == 1 {
		return arr
	}
	temp := arr[0]
	for i := 1; i < len(arr); i++ {
		arr[i-1] = arr[i]
	}
	arr[len(arr)-1] = temp
	return arr
}
