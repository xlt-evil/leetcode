package question

import "sort"

func subsetsWithDup(nums []int) [][]int {
	result := make([][]int, 0)
	r := make([]int, 0)
	sort.Ints(nums)
	back2(nums, &result, &r, 0)
	return result
}

func back2(nums []int, result *[][]int, r *[]int, start int) {
	cur := make([]int, len(*r))
	copy(cur, *r) // copy 是目标文件多大复制多少
	*result = append(*result, cur)
	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		*r = append(*r, nums[i])
		back1(nums, result, r, i+1)
		*r = (*r)[:len(*r)-1]
	}
}
