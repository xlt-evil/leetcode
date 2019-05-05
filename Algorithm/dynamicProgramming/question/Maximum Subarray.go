package question

func MaxSubArray(nums []int) int {
	length := len(nums)
	if length < 1 {
		return 0
	}
	sum := nums[0] // 记录当前最大子串值
	res := nums[0] // 记录最大子串值
	for i := 1; i < length; i++ {
		sum = max5(sum+nums[i], nums[i])
		res = max5(res, sum)
	}
	return res
}

func max5(a, b int) int {
	if a > b {
		return a
	}
	return b
}
