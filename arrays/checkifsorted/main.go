package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{5, 4, 6, 7, 8}

	fmt.Println("Is array ", arr1, " sorted? ", issorted(arr1))
	fmt.Println("Is array ", arr2, " sorted? ", issorted(arr2))
}

func issorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1] < arr[i] {
			return false
		}
	}
	return true
}
