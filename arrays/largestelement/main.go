package main

import "fmt"

func main() {
	var arr1 = []int{2, 5, 1, 3, 0}
	var arr2 = []int{8, 10, 5, 7, 9}

	fmt.Printf("The largest element in arr %v\n", arr1)
	fmt.Println(largestelement(arr1))
	fmt.Printf("The largest element in arr %v\n", arr2)
	fmt.Println(largestelement(arr2))

}

func largestelement(arr []int) int {
	element := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > element {
			element = arr[i]
		}
	}
	return element
}
