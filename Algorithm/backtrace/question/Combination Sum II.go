package question

import "sort"

func CombinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := make([][]int, 0)
	r := make([]int, 0)
	length := len(candidates)
	use := make([]bool, length)
	back4(candidates, &result, &r, length, target, &use, 0)
	return result
}

func back4(nums []int, result *[][]int, r *[]int, length int, target int, use *[]bool, start int) {
	if target == 0 {
		cur := make([]int, len(*r))
		copy(cur, *r)
		*result = append(*result, cur)
		return
	} else if target < 0 {
		return
	}
	for i := start; i < length; i++ {
		if (*use)[i] || i > 0 && nums[i] == nums[i-1] && !(*use)[i-1] || nums[i] > target {
			continue
		}
		(*use)[i] = true
		*r = append(*r, nums[i])
		back4(nums, result, r, length, target-nums[i], use, i+1)
		(*use)[i] = false
		*r = (*r)[:len(*r)-1]
	}
}
