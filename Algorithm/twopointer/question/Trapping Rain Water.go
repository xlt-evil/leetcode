package question

func Trap(height []int) int {
	length := len(height)
	if length < 2 {
		return 0
	}
	left, right := 0, length-1
	area := 0
	maxleft, maxright := 0, 0
	for left <= right {
		if height[left] <= height[right] { // 水从低处流
			if height[left] >= maxleft { //更新存水区域
				maxleft = height[left]
			} else {
				area += maxleft - height[left]
			}
			left++
		} else {
			if height[right] >= maxright { // 更新水的区域
				maxright = height[right]
			} else {
				area += maxright - height[right]
			}
			right--
		}

	}
	return area
}
