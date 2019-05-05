package question

func MaxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	min := prices[0] // 最低买入
	max := prices[1] // 最高卖出
	res := max - min
	for i := 1; i < n; i++ {
		if prices[i] > max {
			max = prices[i]
		}
		if prices[i] < min && i+1 < n {
			min = prices[i]
			max = prices[i+1]
		}
		res = max4(max-min, res)
	}
	return max4(res, 0)
}

func max4(a, b int) int {
	if a > b {
		return a
	}
	return b
}
