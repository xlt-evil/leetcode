package question

func MySqrt(x int) int {
	head, tail := 0, x
	for head < tail {
		middle := (head + tail) / 2
		num := middle * middle
		if num == x {
			return middle
		}
		if num > x {
			tail = middle - 1
			continue
		}
		if num < x {
			head = middle + 1
		}
	}
	if head*head > x {
		return head - 1
	}
	return head
}
