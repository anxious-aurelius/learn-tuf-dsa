package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 4, 7, 7, 5}
	arr2 := []int{1}

	fmt.Println("Second largest number in ", arr1)
	fmt.Println(secondlargestnumber(arr1))
	fmt.Println("Second largest number in ", arr2)
	fmt.Println(secondlargestnumber(arr2))

}

func secondlargestnumber(arr []int) int {
	if len(arr) < 2 {
		return -1
	}
	largest := arr[0]
	secondlargest := arr[0]

	for _, val := range arr {
		if val > largest {
			secondlargest = largest
			largest = val
		} else if val > secondlargest && val != largest {
			secondlargest = val
		}
	}
	return secondlargest
}
