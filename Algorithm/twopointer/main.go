package main

import (
	"fmt"
	"leecode/twopointer/question"
)

func main() {
	fmt.Println("双指针算法")
	fmt.Println("1.盛水的最大容器leetcode-11")
	fmt.Println(question.MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println("2.找到一组3元组，与给定目标最近leetcode-16")
	fmt.Println(question.ThreeSumClosest([]int{-1, 2, 1, -4}, 1))
	fmt.Println("3.找到3元组使得其总和为0，且3元组集合不能重复leetcode-15")
	fmt.Println(question.ThreeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println("4.找到重复元素在时间O(n^2)之内，空间O(1)不能修改数组leetcode-287")
	fmt.Println(question.FindDuplicate([]int{2, 5, 9, 6, 9, 3, 8, 9, 7, 1}))
	fmt.Println("5.颜色排序，leetcode-75")
	fmt.Println(question.SortColors([]int{3, 2, 1, 5, 3}))
	fmt.Println("6.诱捕雨水,leetcode-42")
	fmt.Println(question.Trap([]int{4, 2, 3}))
	fmt.Println("7.找到目标的和排序好的数组中.leedcode-167")
	fmt.Println(question.TwoSum([]int{-1, 0}, -1))
}
