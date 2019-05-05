package question

func permute(nums []int) [][]int {
	length := len(nums)
	result := make([][]int, 0)
	r := make([]int, 0)
	record := map[int]bool{}
	back(nums, &result, &r, length, record)
	return result
}

//回溯法
func back(nums []int, result *[][]int, r *[]int, length int, record map[int]bool) {
	if len(*r) == length {
		cur := make([]int, length)
		copy(cur, *r)
		*result = append(*result, cur)
		return
	}
	for i, _ := range nums {
		if _, ok := record[nums[i]]; ok {
			continue
		}
		*r = append(*r, nums[i])
		record[nums[i]] = true
		back(nums, result, r, length, record)
		delete(record, nums[i])
		*r = (*r)[:len(*r)-1]
	}
}
