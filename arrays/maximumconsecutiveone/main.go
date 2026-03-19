package main

import "fmt"

func main() {
	arr := []int{1, 1, 0, 1, 1, 1}
	fmt.Println(maxones(arr))
}

func maxones(arr []int) int {
	streak := 0
	maxStreak := 0
	for _, val := range arr {
		if val == 1 {
			streak++
			if streak > maxStreak {
				maxStreak = streak
			}
		} else {
			streak = 0
		}

	}
	return maxStreak
}
