package question

//头尾双指针
func MaxArea(height []int) int {
	xlength := len(height)
	if xlength == 0 {
		return 0
	}
	low, high := 0, xlength-1
	area, ylength := 0, 0
	xlength--
	for low < high {
		if height[low] > height[high] {
			ylength = height[high]
			high--
		} else if height[low] <= height[high] {
			ylength = height[low]
			low++
		}
		if area < ylength*(xlength) {
			area = ylength * (xlength)
		}
		xlength--
	}
	return area
}
