package main

import "fmt"

func strStr(haystack string, needle string) int {
	var index = -1
	needLength := len(needle)
	haylength := len(haystack)
	if haylength == 0 {
		index = -1
	}
	if needLength == 0 {
		index = 0
	} else {
		for i := 0; i <= haylength-needLength; i++ {
			str := haystack[i : needLength+i]
			if needle == str {
				index = i
				break
			}
		}
	}

	return index
}

func main() {
	fmt.Println(strStr("mississippi", "issipi"))
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("babbbbbabb", "bbab"))
	fmt.Println(strStr("bbbbababbbaabbba", "abb"))
	fmt.Println(strStr("bbaa", "aab"))
}
