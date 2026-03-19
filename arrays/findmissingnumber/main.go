package main

import "fmt"

func main() {
	arr := []int{6, 2, 4, 5, 3, 7, 8}
	fmt.Println(sumfindmissing(arr))
}

func sumfindmissing(arr []int) int {
	n := len(arr) + 1
	sum := n * (n + 1) / 2
	for _, val := range arr {
		sum -= val
	}
	return sum
}

//func hashmapfindmissing(arr []int) int {
//	freq := make(map[int]int)
//	for i := 1; i <= len(arr)+1; i++ {
//		freq[i] = 0
//	}
//
//	for _, val := range arr {
//		freq[val] = 1
//	}
//
//	for key, val := range freq {
//		if val == 0 {
//			return key
//		}
//	}
//	return -1
//}

// using linear searching technique for each expected element
//func linearfindmissing(arr []int) int {
//	var num int
//	for num = 1; num <= len(arr)+1; num++ {
//		numFound := false
//		for _, value := range arr {
//			if value == num {
//				numFound = true
//			}
//		}
//		if !numFound {
//			break
//		}
//	}
//	return num
//}
