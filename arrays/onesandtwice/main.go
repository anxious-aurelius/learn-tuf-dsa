package main

import "fmt"

func main() {
	arr := []int{4, 1, 2, 1, 2, 6, 4}
	fmt.Println(appearsOnes(arr))
}

func appearsOnes(arr []int) int {
	result := 0
	for _, val := range arr {
		result ^= val
	}
	return result
}

// function uses hashmap to find the number occurring ones
//func appearsOnes(arr []int) int {
//	freq := make(map[int]int)
//	for _, val := range arr {
//
//		freq[val]++
//	}
//
//	for key, val := range freq {
//		if val == 1 {
//			return key
//		}
//	}
//
//	return -1
//}
