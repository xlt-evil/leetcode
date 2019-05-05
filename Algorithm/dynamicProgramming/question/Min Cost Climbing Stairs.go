package question

func MinCostClimbingStairs(cost []int) int {
	n := len(cost) // n 为攀爬的高度
	if n < 2 {
		return 0
	}
	p1, p2 := cost[0], cost[1] // 设置两个小人初始在第一台阶和第二台阶上
	pay := 0                   // 代价
	for i := 2; i < n; i++ {
		pay = min3(p1, p2) + cost[i] // 每一步选择最小的代价前进
		p1 = p2
		p2 = pay
	}
	return min3(p1, p2)
}

func min3(a, b int) int {
	if a > b {
		return b
	}
	return a
}
