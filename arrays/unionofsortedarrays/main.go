package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{2, 4, 6, 8, 10}

	fmt.Println(unionarray(arr1, arr2))
}

func unionarray(nums1 []int, nums2 []int) []int {
	idx1 := 0
	idx2 := 0
	result := make([]int, 0, len(nums1)+len(nums2))

	if nums1[0] > nums2[0] {
		result = append(result, nums2[0])
	} else {
		result = append(result, nums1[0])
	}

	for idx1 < len(nums1) && idx2 < len(nums2) {
		if nums1[idx1] < nums2[idx2] {
			if result[len(result)-1] != nums1[idx1] {
				result = append(result, nums1[idx1])
			}
			idx1++
		} else if nums1[idx1] > nums2[idx2] {
			if result[len(result)-1] != nums2[idx2] {
				result = append(result, nums2[idx2])
			}
			idx2++
		} else {
			idx2++
		}
	}
	for idx1 < len(nums1) {
		result = append(result, nums1[idx1])
		idx1++
	}
	for idx2 < len(nums2) {
		result = append(result, nums2[idx2])
		idx2++
	}
	return result
}
