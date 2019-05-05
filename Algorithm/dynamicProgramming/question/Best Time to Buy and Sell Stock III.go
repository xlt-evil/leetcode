package question

//利用状态机的方式实现DP

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

func MaxProfit3(prices []int) int {
	length := len(prices)
	if length < 1 {
		return 0
	}
	s1, s2, s3, s4 := -prices[0], INT_MIN, INT_MIN, INT_MIN
	for i := 1; i < length; i++ {
		s1 = max7(s1, -prices[i])
		s2 = max7(s2, s1+prices[i])
		s3 = max7(s3, s2-prices[i])
		s4 = max7(s4, s3+prices[i])
	}
	return max7(0, s4)
}

func max7(a, b int) int {
	if a > b {
		return a
	}
	return b
}
