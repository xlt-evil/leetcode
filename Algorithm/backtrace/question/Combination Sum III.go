package question

func CombinationSum3(k int, n int) [][]int {
	result := make([][]int, 0)
	temp := make([]int, 0)
	this := 0
	back7(&result, k, n, &temp, this, 1)
	return result
}

func back7(result *[][]int, k, n int, temp *[]int, this int, index int) {
	if len(*temp) == k && this == n {
		cur := make([]int, len(*temp))
		copy(cur, *temp)
		*result = append(*result, cur)
		return
	}
	for i := index; i < n; i++ {
		if this+i <= n {
			*temp = append(*temp, i)
			this += i
			back7(result, k, n, temp, this, i+1)
			*temp = (*temp)[:len(*temp)-1]
			this -= i
		}
	}
}
