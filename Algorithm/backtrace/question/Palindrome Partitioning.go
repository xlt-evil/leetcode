package question

func partition(s string) [][]string {
	result := make([][]string, 0)
	r := make([]string, 0)
	length := len(s)
	back5(s, &result, &r, 0, length)
	return result
}

func back5(s string, result *[][]string, r *[]string, start int, length int) {
	if start == length {
		cur := make([]string, len(*r))
		copy(cur, *r)
		*result = append(*result, cur)
		return
	}

	for i := start; i < length; i++ {
		if isPalindrome(start, i, s) {
			*r = append(*r, s[start:i+1])
			back5(s, result, r, i+1, length)
			*r = (*r)[:len(*r)-1]
		}
	}
}

func isPalindrome(start, end int, s string) bool {
	low, high := start, end
	for low < high {
		if s[low] != s[high] {
			return false
		}
		low++
		high--
	}
	return true
}
