package question

func IsUgly(num int) bool {
	if num == 1 || num == 0 {
		return false
	}
	list := []int{5, 3, 2}
	answer := false
	back10(num, list, 0, 1, num, &answer)
	return answer
}

func back10(num int, list []int, index, result int, target int, answer *bool) {

	if result == target {
		*answer = true
		return
	}
	for i := index; i < len(list); i++ {
		if num%list[i] == 0 {
			result *= list[i]
			back10(num/list[i], list, i, result, target, answer)
			if *answer {
				return
			}
			result /= list[i]
		}
	}
}
