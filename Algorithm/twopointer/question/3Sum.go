package question

import "sort"

func ThreeSum(nums []int) [][]int {
	length := len(nums)
	if length < 3 {
		return nil
	}
	sort.Ints(nums)
	result := make([][]int, 0)
	record := map[[3]int]bool{}
	for i := 1; i <= length-2; i++ {
		low, high := i-1, i+1
		for low >= 0 && high <= length-1 {
			sum := nums[i] + nums[low] + nums[high]
			if sum == 0 {
				temp := [3]int{nums[low], nums[i], nums[high]}
				low--
				high++
				if _, ok := record[temp]; ok {
					continue
				}
				result = append(result, temp[:])
				record[temp] = true
				continue
			}
			if sum > 0 {
				low--
			} else {
				high++
			}

		}
	}
	return result
}
