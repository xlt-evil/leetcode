package question

import "sort"

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}
	sort.Ints(nums)
	result := make([][]int, 0)
	r := make([]int, 0)
	length := len(nums)
	use := make([]bool, length)
	sum := 0
	back6(nums, &result, &r, target, 0, &use, length, sum)
	return result
}

func back6(num []int, result *[][]int, r *[]int, target int, start int, use *[]bool, length int, sum int) {
	if len(*r) == 4 {
		if target == sum {
			cur := make([]int, 4)
			copy(cur, *r)
			*result = append(*result, cur)
		}
		return
	}

	for i := start; i < length; i++ {
		if (*use)[i] || i > 0 && num[i] == num[i-1] && !(*use)[i-1] { // 求子集切不重复的基本写法
			continue
		}
		*r = append(*r, num[i])
		(*use)[i] = true
		back6(num, result, r, target, i+1, use, length, sum+num[i])
		*r = (*r)[:len(*r)-1]
		(*use)[i] = false
	}
}
