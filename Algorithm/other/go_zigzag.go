package main

import "fmt"

//ZigZag 转换
//1   5   9
//2 4 6 8 10
//3   7   11
//特性：增删行都会发生偶数次的差距 ，头尾两行相差数一样 ，之间行是由两个等差组成的数列左边数字到右边斜线的数字是以2n的由上之下逐行递减，左边的斜线到右边由上之下逐行递减，每行的两个等差相加都相当,
//每行的开始都是按原字符串顺序下去的
func convert(s string, numRows int) string {
	result := ""
	if len(s) <= 2 || numRows == 1 {
		return s
	}
	r := make([]rune, 0)
	for _, v := range s {
		r = append(r, v)
	}
	length := len(r)
	for i := 0; i < numRows; i++ {
		d1 := (numRows - i - 1) * 2
		d2 := i * 2
		index := i
		if index < length {
			result += string(r[index])
		}
		for {
			index += d1
			if index >= length {
				break
			}
			if d1 != 0 {
				result += string(r[index])
			}
			index += d2
			if index >= length {
				break
			}
			if d2 != 0 { //防止重复
				result += string(r[index])
			}
		}
	}
	return result
}

func main() {
	fmt.Println(convert("AB", 1))
}
