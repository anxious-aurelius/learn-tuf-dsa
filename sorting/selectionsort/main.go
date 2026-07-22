package main

import "fmt"

func main(){
	arr := []int{4,1,5,2,5,12,0}

	fmt.Println("Array before sorting : ", arr)
	selectionSort(arr)
	fmt.Println("Array after sorting : ", arr)

}

func selectionSort( arr []int) {

	for i := 0; i < len(arr) - 1; i++ {
		minIndex := i
		for j :=  i + 1 ; j < len(arr); j++ {
			if arr[j] < arr[minIndex]{
				minIndex = j
			}
		}
		if minIndex != i {
			arr[minIndex], arr[i] = arr[i], arr[minIndex]
		}
	}

}