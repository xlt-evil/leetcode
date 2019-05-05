package question

import "sort"

func Intersect(nums1 []int, nums2 []int) []int {
	record := map[int]int{}
	for i, _ := range nums1 {
		_, ok := record[nums1[i]]
		if ok {
			record[nums1[i]]++
		} else {
			record[nums1[i]] = 1
		}
	}
	result := []int{}
	for i, _ := range nums2 {
		if v, ok := record[nums2[i]]; ok {
			if v != 0 {
				result = append(result, nums2[i])
				record[nums2[i]]--
			}
		}
	}
	return result
}
