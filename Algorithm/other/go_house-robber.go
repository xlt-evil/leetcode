package main

import "fmt"

func rob(nums []int) int {
	length := len(nums)
	if length <= 2 {
		if length == 1 {
			return nums[0]
		}
		if length == 0 {
			return 0
		}
		if length == 2 {
			if nums[0] > nums[1] {
				return nums[0]
			} else {
				return nums[1]
			}
		}
	}
	i1, j1 := 0, length-1
	sum1, sum2 := 0, 0
	i2, j2 := 0, length-2
	for i1 < j1 {
		sum1 += nums[i1] + nums[j1]
		i1 += 2
		j1 -= 2
	}

	for i2 < j2 {
		sum2 += nums[i2] + nums[j2]
		i2 += 2
		j2 -= 2
	}
	if sum1 > sum2 {
		return sum1
	}
	return sum2
}

func rotate(nums []int, k int) {
	length := len(nums)
	for k > 0 {
		temp := nums[length-1]
		for i := length - 2; i >= 0; i-- {
			nums[i+1] = nums[i]
		}
		nums[0] = temp
		k--
	}
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(s, 3)
	fmt.Println(rob(s))
	fmt.Println(s)
}
