package question

func SearchInsert(nums []int, target int) int {
	length := len(nums) - 1
	head, fail := 0, length
	for head <= fail {
		middle := (head + fail) / 2
		if nums[middle] == target {
			return middle
		}
		if nums[middle] < target {
			head = middle + 1
		} else if nums[middle] > target {
			fail = middle - 1
		}
	}
	return head
}
