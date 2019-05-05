package main

import (
	"fmt"
	"strconv"
)

func myAtoi(str string) int {
	r := make([]rune, 0)
	for _, v := range str {
		r = append(r, v)
	}
	flag := true //第一个符号判断
	state := 1   //默认整数
	result := ""
	length := len(r)
	for i := 0; i < length; i++ {
		if flag {
			if string(r[i]) == " " {
				continue
			}
			if r[i] >= 48 && r[i] <= 57 || r[i] == 43 || r[i] == 45 {
				flag = false
				if r[i] >= 48 && r[i] <= 57 {
					result += string(r[i])
				} else if r[i] == 45 {
					state = -1
				}
			} else {
				return 0
			}
		} else {
			if r[i] >= 48 && r[i] <= 57 {
				result += string(r[i])
			} else {
				break
			}
		}
	}
	myAtoi, _ := strconv.Atoi(result)
	if state > 1 {
		if myAtoi > 1<<31-1 {
			myAtoi = 2147483647
		}
	} else {
		if myAtoi > 1<<31 {
			myAtoi = state * 2147483648
		} else {
			myAtoi = state * myAtoi
		}
	}
	return myAtoi
}

func main() {
	fmt.Println(myAtoi("2147483648"))
}
