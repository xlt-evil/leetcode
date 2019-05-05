package question

func maxProduct(nums []int) int {
	length := len(nums)
	if length < 1 {
		return 0
	}
	f := make([]int, length) // 保存每一步的乘积尽可能最大
	g := make([]int, length) // 保存每一步的乘积尽可能最小
	f[0] = nums[0]
	g[0] = nums[0]
	res := nums[0]
	for i := 1; i < length; i++ {
		f[i] = max(max(f[i-1]*nums[i], g[i-1]*nums[i]), nums[i])
		g[i] = mins(mins(f[i-1]*nums[i], g[i-1]*nums[i]), nums[i])
		res = max(f[i], res)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func mins(a, b int) int {
	if a < b {
		return a
	}
	return b
}
