package question

func MinimumTotal(triangle [][]int) int {
	length := len(triangle)
	if length < 1 {
		return 0
	}
	for i := 1; i < length; i++ {
		widths := len(triangle[i])
		for j, _ := range triangle[i] {
			if j == 0 {
				triangle[i][j] += triangle[i-1][0]
				continue
			}
			if j == widths-1 {
				triangle[i][j] += triangle[i-1][j-1]
				continue
			}
			min := triangle[i-1][j-1]
			if min > triangle[i-1][j] {
				min = triangle[i-1][j]
			}
			triangle[i][j] += min
		}
	}
	min := triangle[length-1][0]
	for _, v := range triangle[length-1] {
		if v < min {
			min = v
		}
	}
	return min
}

//自底向上
func minimumTotal(triangle [][]int) int {
	r := len(triangle)
	for i := r - 2; i >= 0; i-- {
		for j := 0; j < i+1; j++ {
			triangle[i][j] += min(triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	return triangle[0][0]
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
