package main

import (
	"fmt"
	"leecode/binarysearch/question"
)

func main() {
	fmt.Println("二分搜索")
	fmt.Println("寻找目标位置，leetcode-35")
	fmt.Println(question.SearchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println("查找平分根，leetcode-69")
	fmt.Println(question.MySqrt(0))
	fmt.Println(question.Intersect([]int{1, 2, 2, 1}, []int{2, 2}))
}
