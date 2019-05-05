package main

import "fmt"

func removeDuplicates(nums []int) int {
	num := map[int]bool{}
	length := 0
	for i, _ := range nums {
		if _, ok := num[nums[i]]; ok {
			continue
		}
		num[nums[i]] = true
		nums[length] = nums[i]
		length++
	}
	return length
}

func main() {
	s := []int{1, 2, 3, 2}
	removeDuplicates(s)
	fmt.Println(s)
}
