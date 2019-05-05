package question

type NumArray struct {
	Sum []int
}

func Constructor(nums []int) NumArray {
	length := len(nums)
	start := nums[0]
	for i := 1; i < length; i++ {
		nums[i] += start
		start = nums[i]
	}
	return NumArray{Sum: nums}
}

func (this *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		return this.Sum[j]
	}
	return this.Sum[j] - this.Sum[i-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
