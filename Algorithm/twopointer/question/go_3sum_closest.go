package question

import (
	"math"
	"sort"
)

//从某个点向两头双指针
func ThreeSumClosest(nums []int, target int) int {
	length := len(nums)
	if length < 3 {
		return 0
	}
	sort.Ints(nums)
	result := nums[0] + nums[1] + nums[2]
	//以当前元素为核心找另外两个数
	for i := length - 1; i >= 0; i-- {
		start := i + 1
		end := length - 1
		for start < end {
			//当前的和
			sum := nums[i] + nums[start] + nums[end]
			if sum == target {
				return sum
			}
			//距离更小的越接近target
			if math.Abs(float64(target)-float64(sum)) < math.Abs(float64(target)-float64(result)) {
				result = sum
			}
			if sum > target {
				end--
			} else {
				start++
			}
		}
	}
	return result
}
