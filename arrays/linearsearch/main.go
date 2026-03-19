package main

import "fmt"

func main() {
	arr := []int{1, 2, 34, 5, 12, 19, 2, 12, 53}
	fmt.Println(linearsearch(arr, 19))

}

func linearsearch(nums []int, num int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == num {
			return i
		}
	}
	return -1
}
