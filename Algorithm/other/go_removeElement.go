package main

// 移除重复下标
func removeElement(nums []int, val int) int {
	length := 0
	for i, _ := range nums {
		if nums[i] == val {
			continue
		}
		nums[length] = nums[i]
	}
	return length
}
