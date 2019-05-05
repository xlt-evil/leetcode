package question

func TwoSum(numbers []int, target int) []int {
	length := len(numbers)
	if length < 1 {
		return nil
	}
	head, tail := 0, length-1
	for head < tail {
		if numbers[tail]+numbers[head] == target {
			return []int{head, tail}
		}
		if numbers[head]+numbers[tail] < target {
			head++
			continue
		}
		tail--
	}
	return nil
}
