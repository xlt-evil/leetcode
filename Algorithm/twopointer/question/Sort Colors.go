package question

//冒泡
func SortColors(nums []int) []int {
	length := len(nums)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

// 快排 -- 双指针
func QuickSort(nums []int) {
	length := len(nums)
	if length < 1 {
		return
	}
	left, right := 0, length-1
	temp := nums[left] // 观察点
	for left <= right {
		if nums[right] < temp { //找到小于基准点的数
			for left != right && nums[left] <= temp {
				left++
			}
			nums[left], nums[right] = nums[right], nums[left]
		}
		right--
	}
	nums[0], nums[left] = nums[left], temp //left 最后的坐标一定是小于基准数的，一次循环结束temp 的左边一定全部小于temp但不一定有序,右边则是全部大于
	QuickSort(nums[:left])
	QuickSort(nums[left+1:])
}
