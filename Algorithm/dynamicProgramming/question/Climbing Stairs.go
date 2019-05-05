package question

// 利用dp 求解
// 最优子结构
// 边界
// 状态转移公式
func ClimbStairs(n int) int {
	res := 0
	tem := 0
	for i := 0; i < n; i++ {
		if i == 0 {
			res = 1
			tem = res
		} else {
			res += tem      // 下一次的循环是i - 1
			tem = res - tem //保存i - 2 的走法
		}

	}
	return res
}
