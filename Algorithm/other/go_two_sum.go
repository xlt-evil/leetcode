package main

import "fmt"

//两个数之和
func main() {
	num := []int{2, 7, 11, 15}
	fmt.Println(TwoSum(9, num))
}

func TwoSum(target int, arr []int) (result []int) {
	m := map[int]int{}
	for i, v := range arr {
		if j, ok := m[target-v]; ok {
			result = append(result, j)
			result = append(result, i)
		} else {
			m[v] = i
		}
	}
	return
}
