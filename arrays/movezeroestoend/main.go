package main

import "fmt"

func main() {
	arr := []int{0, 1, 0, 3, 12}
	fmt.Println(movetoend(arr))
}

func movetoend(arr []int) []int {
	zindex := -1
	for index, value := range arr {
		if value == 0 {
			zindex = index
			break
		}
	}
	if zindex != -1 {
		i := zindex + 1
		for i < len(arr) {
			if arr[i] == 0 {
				i++
			} else {
				arr[zindex] = arr[i]
				arr[i] = 0
				zindex++
			}
		}
	}
	return arr
}

//func movetoend(arr []int, num int) []int {
//	var wasSwapped bool = false
//	for i := 0; i < len(arr); i++ {
//		wasSwapped = false
//		for j := 0; j < len(arr)-1; j++ {
//			if arr[j] == num {
//				arr[j] = arr[j+1]
//				arr[j+1] = num
//				wasSwapped = true
//			}
//		}
//		if !wasSwapped {
//			break
//		}
//	}
//	return arr
//}
