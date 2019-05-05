package main

import (
	"fmt"
)

type P struct {
	Val   int
	Index int
}

func jump(nums []int) int {
	length := len(nums)
	if length == 0 || length == 1 {
		return 0
	}
	minStep := 0
	Q := []P{P{nums[0], 0}}
	end := length - 1
	for len(Q) > 0 {
		minStep++
		cur := Q[0]
		Q = Q[1:]
		if cur.Index == end {
			return minStep
		}
		if cur.Val+cur.Index >= end {
			return minStep
		}
		max := P{}
		maxdistance := 0
		for i := cur.Index + 1; i < cur.Index+cur.Val+1 && i < length; i++ {
			if maxdistance < nums[i]+i {
				if maxdistance >= length {
					minStep++
					return minStep
				}
				max.Val = nums[i]
				max.Index = i
				maxdistance = nums[i] + i
			}
		}
		Q = append(Q, max)
	}
	return -1
}

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}), 2)
	fmt.Println(jump([]int{1, 2, 3}), 2)
	fmt.Println(jump([]int{1, 2}), 1)
	fmt.Println(jump([]int{2, 1}), 1)
}
