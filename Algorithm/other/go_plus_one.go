package main

import ()

func main() {

}

//[7,2,8,5,0,9,1,2,9,5,3,6,6,7,3,2,8,4,3,7,9,5,7,7,4,7,4,9,4,7,0,1,1,1,7,4,0,0,6]
func plusOne(digits []int) []int {
	length := len(digits)
	count := 0
	for j := length - 1; j >= 0; j-- {
		count = digits[j] + 1
		if count >= 10 {
			digits[j] = 0
			continue
		}
		digits[j] += 1
		count = 0
		break
	}
	//Indicates that the highest position has a carry.
	if count != 0 {
		digits[0] = 0
		temp := []int{}
		temp = append(temp, 1)
		temp = append(temp, digits...)
		return temp
	}
	return digits
}
