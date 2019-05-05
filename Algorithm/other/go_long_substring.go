package main

import (
	"fmt"
)

//没有重复字符的最长子串
//滑动窗口思想
func lengthOfLongestSubstring(s string) int {
	byteArr := s[0:]
	charArr := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		charArr[i] = int(byteArr[i])
	}

	var length int

	dupMap := make(map[int]int)

	var slow, fast int

	for fast < len(s) {
		if len(s)-slow < length {
			break
		}

		idx, ok := dupMap[charArr[fast]]

		if ok && idx >= slow {
			if length < fast-slow {
				length = fast - slow
			}
			slow = idx + 1
		}

		if length < fast+1-slow { //新子串的长度
			length = fast + 1 - slow
		}

		dupMap[charArr[fast]] = fast //保存最新的字符的滑动坐标
		fast++
	}
	return length
}

func main() {
	str := "abcabcbb"
	fmt.Println("实际长度", lengthOfLongestSubstring(str), "期望长度：3")
	str = "bbbbb"
	fmt.Println("实际长度", lengthOfLongestSubstring(str), "期望长度：1")
	str = "pwwkew"
	fmt.Println("实际长度", lengthOfLongestSubstring(str), "期望长度：3")
}
