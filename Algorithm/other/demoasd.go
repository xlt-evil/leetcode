package main

import "fmt"

func Remove(nums []int) int {
	index := 0
	length := len(nums)
	for i := 1; i < length; i++ {
		if nums[i] != nums[index] {
			index++
			nums[index] = nums[i]
		}
	}
	return index + 1
}

func main() {
	fmt.Println(Remove([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}
