package main

import (
	"fmt"
	"sort"
)

//找到连个排序数组的中位数

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	arrays := append(nums1, nums2...)
	length := len(arrays)
	answer := 0.0
	sort.Ints(arrays)
	if length%2 == 0 {
		id1 := length / 2
		id2 := id1 - 1
		answer = float64(arrays[id1]+arrays[id2]) / 2
	} else {
		answer = float64(arrays[length/2])
	}

	return answer
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))

}
