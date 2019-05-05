package question

func PredictTheWinner(nums []int) bool {
	length := len(nums)
	if length < 2 {
		return true
	}
	dp := make([]int, length)
	for i := length - 1; i >= 0; i-- {
		for j := i; j < length; j++ {
			if i == j {
				dp[i] = nums[i]
			} else {
				dp[j] = max6(nums[i]-dp[j], nums[j]-dp[j-1])
			}
		}
	}
	return dp[length-1] >= 0
}

func max6(a, b int) int {
	if a > b {
		return a
	}
	return b
}
