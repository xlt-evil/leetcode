package main

import (
	"fmt"
	"strings"
)

// 寻找前缀
func longestCommonPrefix(strs []string) string {
	length := len(strs)
	if strs == nil || length == 0 {
		return ""
	}
	prefix := strs[0]
	i := 1
	for i < length {
		for strings.Index(strs[i], prefix) != 0 {
			if len(prefix) == 0 {
				break
			}
			prefix = prefix[:len(prefix)-1]
		}
		i++
	}
	return prefix
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"c", "acc", "ccc"}))
	fmt.Println(strings.Index("134", "3"))
}
