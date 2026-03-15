package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 0, 1, 4, 0, 0, 0, 5, 1, 2, 0, 1}
	fmt.Println(movetoend(arr, 0))
}

func movetoend(arr []int, num int) []int {
	var wasSwapped bool = false
	for i := 0; i < len(arr); i++ {
		wasSwapped = false
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] == num {
				arr[j] = arr[j+1]
				arr[j+1] = num
				wasSwapped = true
			}
		}
		if !wasSwapped {
			break
		}
	}
	return arr
}
