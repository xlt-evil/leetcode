package question

import "sort"

func PermuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	r := make([]int, 0)
	length := len(nums)
	use := make([]bool, length)
	back3(nums, &result, &r, length, &use)
	return result
}

func back3(nums []int, result *[][]int, r *[]int, length int, use *[]bool) {
	if len(*r) == length {
		cur := make([]int, length)
		copy(cur, *r)
		*result = append(*result, cur)
		return
	}
	for i := 0; i < length; i++ {
		if (*use)[i] || i > 0 && nums[i-1] == nums[i] && !(*use)[i-1] { // use 跳过当前已添加元素，并且判断最新元素是否跟之前元素相等，如果最新元素跟之前元素相等，但是之前元素没有使用则跳过，因为选择之前的元素会是同样的结果
			continue
		}
		(*use)[i] = true
		*r = append(*r, nums[i])
		back3(nums, result, r, length, use)
		(*use)[i] = false
		*r = (*r)[:len(*r)-1]
	}
}
