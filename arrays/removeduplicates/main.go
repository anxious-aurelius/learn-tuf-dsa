package main

import "fmt"

func main() {
	arr1 := []int{1, 1, 2, 2, 2, 3, 3, 4, 4, 4, 4, 4, 4, 5, 5, 6, 7}
	fmt.Println(removeduplicates(arr1))
}

func removeduplicates(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	uniqueindex := 0
	for _, value := range arr {
		if value == arr[uniqueindex] {
			continue
		} else {
			uniqueindex++
			arr[uniqueindex] = value
		}
	}
	return uniqueindex + 1

}
