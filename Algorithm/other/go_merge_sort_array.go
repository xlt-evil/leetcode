package main

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n <= 0 {
		return
	}

	if len(nums1) < m+n {
		panic("len nums1 not enough")
	}

	var i = m - 1
	var j = n - 1
	var toFill = m + n - 1
	for ; i >= 0 && j >= 0; toFill-- {
		if nums1[i] > nums2[j] {
			nums1[toFill] = nums1[i]
			i--
		} else {
			nums1[toFill] = nums2[j]
			j--
		}
	}
	for ; j >= 0; j = j - 1 {
		nums1[j] = nums2[j]
	}
}

func main() {
	nums11 := []int{4, 5, 6, 0, 0, 0}
	nums12 := []int{1, 2, 3}
	merge(nums11, 3, nums12, 3)
	fmt.Println(nums11)
}
