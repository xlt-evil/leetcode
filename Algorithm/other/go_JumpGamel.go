package main

import "fmt"

func main() {
	nums := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(nums))
}

type Point struct {
	X int
	Y int
}

func canJump(nums []int) bool {
	length := len(nums)
	if length == 0 {
		return true
	}
	Q := []Point{Point{X: 0, Y: nums[0]}}
	fd := Q[0].X + Q[0].Y //farthest index
	for len(Q) > 0 {
		this := Q[0]
		Q = Q[1:]
		if this.X+this.Y >= length-1 {
			return true
		}
		for i, j := this.X+1, 0; j < this.Y; i, j = i+1, j+1 {
			if i < length {
				if i+nums[i] >= length-1 {
					return true
				}
				if i+nums[i] >= fd {
					p := Point{X: i, Y: nums[i]}
					Q = append(Q, p)
					fd = this.X + this.Y
				}

			}
		}
	}
	return false
}
