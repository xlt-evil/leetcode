package main

import (
	"fmt"
	"sort"
)

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	res := make([]int, 0)
	length1 := len(nums1)
	length2 := len(nums2)
	i, j, temp := 0, 0, 0
	for ; i < length1; i++ {
		j = temp
		for ; j < length2; j++ {
			if nums1[i] == nums2[j] {
				temp = j + 1
				res = append(res, nums1[i])
				break
			}
		}
	}
	return res
}

func main() {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
}
