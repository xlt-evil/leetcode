package main

import "fmt"

func numSquares(n int) int {
	if n < 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	count := 0
	perfectNums := []int{}
	for i := 1; i <= n/2; i++ {
		number := i * i
		if i == n {
			return 1
		}
		perfectNums = append(perfectNums, number)
	}
	Q := []int{n}
	for len(Q) > 0 {
		length := len(Q)
		count++
		for i := 0; i < length; i++ {
			cur := Q[0]
			Q = Q[1:]
			for _, v := range perfectNums {
				if cur-v > 0 {
					Q = append(Q, cur-v)
					continue
				}
				if cur-v == 0 {
					return count
				}
				continue
			}
		}
	}
	return count
}

func main() {
	fmt.Println(numSquares(13))
}
