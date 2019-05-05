package question

//二维DP
func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	length := len(obstacleGrid)
	if length < 1 {
		return 0
	}
	width := len(obstacleGrid[0])
	if width < 1 {
		return 0
	}
	DP := make([][]int, length)
	for i, _ := range obstacleGrid {
		DP[i] = make([]int, len(obstacleGrid[i]))
	}
	if obstacleGrid[0][0] == 0 {
		DP[0][0] = 1 //这个点只有一种走法
	}
	for i := 0; i < length; i++ {
		for j, _ := range obstacleGrid[i] {
			if obstacleGrid[i][j] == 1 {
				DP[i][j] = 0
				continue
			}
			if i-1 >= 0 && obstacleGrid[i-1][j] != 1 && j-1 >= 0 && obstacleGrid[i][j-1] != 1 {
				DP[i][j] = DP[i-1][j] + DP[i][j-1]
				continue
			}
			if i-1 >= 0 && obstacleGrid[i-1][j] != 1 {
				DP[i][j] = DP[i-1][j]
				continue
			}
			if j-1 >= 0 && obstacleGrid[i][j-1] != 1 {
				DP[i][j] = DP[i][j-1]
			}
		}
	}
	return DP[length-1][width-1]
}

//一维DP
//左右移动都是单前步数的衍生，对角线才是能
func DP1(obstacleGrid [][]int) int {
	length := len(obstacleGrid)
	if length < 1 {
		return 0
	}
	width := len(obstacleGrid[0])
	if width < 1 {
		return 0
	}
	dp := make([]int, width)
	dp[0] = 1
	for i, _ := range obstacleGrid {
		for j, _ := range obstacleGrid[i] {
			if obstacleGrid[i][j] == 1 { //
				dp[j] = 0
				continue
			}
			if j > 0 {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[width-1]
}
