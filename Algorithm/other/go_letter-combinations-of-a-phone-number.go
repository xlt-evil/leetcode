package main

import (
	"fmt"
	"strings"
)

func letterCombinations(digits string) []string {
	length := len(digits)
	if length < 1 {
		return nil
	}
	buttonword := map[string][]string{"1": nil, "2": []string{"a", "b", "c"}, "3": []string{"d", "e", "f"}, "4": []string{"g", "h", "i"}, "5": []string{"j", "k", "l"}, "6": []string{"m", "n", "o"}, "7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"}, "9": []string{"w", "x", "y", "z"}}
	if length == 1 {
		return buttonword[digits]
	}
	key := strings.Split(digits, "")
	answer := buttonword[key[0]]
	for i, _ := range key {
		if i == 0 {
			continue
		}
		temp, ok := buttonword[key[i]]
		if !ok {
			return nil
		}
		if temp == nil {
			continue
		}
		for _, v1 := range answer {
			answer = answer[1:]
			for _, v2 := range temp {
				answer = append(answer, v1+v2)
			}
		}
	}
	return answer
}

func main() {
	fmt.Println(letterCombinations("232"))
}
