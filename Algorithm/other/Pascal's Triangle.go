package main

func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}
	res := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		list := make([]int, i+1)
		for j := 0; j < i+1; i++ {
			if j == 0 || j == i {
				list[j] = 1
				continue
			}
			list[j] = res[i-1][j-1] + res[i-1][j]
		}
		res[i] = list
	}
	return res
}
