package main

import "fmt"

//反向整数

func reverse(x int) int {
	var result int32
	for x != 0 {
		tail := int32(x % 10) //每次取最后一位数
		var newresult int32   //进一位
		newresult = result*10 + tail
		if (newresult-tail)/10 != result { //如果溢出则不相等
			return 0
		}
		result = newresult
		x = x / 10
	}
	return int(result)
}

func main() {
	fmt.Println(reverse(1534236469))
}
