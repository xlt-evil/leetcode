package question

func Rob(nums []int) int {
	rob := 0
	notrob := 0
	for i := 0; i < len(nums); i++ {
		currob := notrob + nums[i]
		notrob = max1(notrob, rob)
		rob = currob
	}
	return max1(rob, notrob)
}

func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}
