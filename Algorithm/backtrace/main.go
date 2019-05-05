package main

import (
	"fmt"
	"leecode/backtrace/question"
)

func main() {
	fmt.Println("回溯算法")
	fmt.Println(question.Subsets([]int{1, 2, 3}))
	fmt.Println("=======================================")
	fmt.Println(question.PermuteUnique([]int{3, 3, 0, 3}))
	fmt.Println("=======================================")
	fmt.Println(question.CombinationSum2([]int{1, 1, 6}, 8))
	fmt.Println(question.CombinationSum3(3, 7))
	fmt.Println(question.IsUgly(6))
}
