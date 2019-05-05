package question

func MinPathSum(grid [][]int) int {
	m := len(grid)
	if m < 1 {
		return 0
	}
	n := len(grid[0])
	if n < 1 {
		return 0
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, len(grid[i]))
	}
	dp[0][0] = grid[0][0]
	for i, _ := range grid {
		for j, _ := range grid[i] {
			if i > 0 && j > 0 {
				min := dp[i-1][j]
				if min > dp[i][j-1] {
					min = dp[i][j-1]
				}
				dp[i][j] = grid[i][j] + min
				continue
			}
			if i-1 >= 0 {
				dp[i][j] = grid[i][j] + dp[i-1][j]
				continue
			}
			if j-1 >= 0 {
				dp[i][j] += grid[i][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
