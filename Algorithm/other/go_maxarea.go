package main

import "fmt"

//大多数水的容器

func main() {
	heigth := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(heigth))
}

type Points struct {
	X int
	Y int
}

func minint(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxArea(height []int) int {
	res, cur := 0, 0
	for i, j := 0, len(height)-1; i < j; {
		cur = minint(height[i], height[j]) * (j - i)
		if cur > res {
			res = cur
		}
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return res
}
