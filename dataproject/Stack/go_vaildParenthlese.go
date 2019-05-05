package main

import (
	"fmt"
	"strings"
)

//校验符号的正确性
func main() {

	fmt.Println(isValid("([]{})"))
}

type Stack struct {
	str []string
}

func (this *Stack) Push(key string) {
	this.str = append(this.str, key)
}

func (this *Stack) Pop() (key string) {
	length := len(this.str)
	if length != 0 {
		temp := this.str[length-1:]
		key = temp[0]
		this.str = this.str[:length-1]
	}
	return
}

func isValid(s string) bool {
	if s == "" {
		return true
	}
	chars := strings.Split(s, "")
	if len(chars)%2 != 0 {
		return false
	}
	strack := Stack{}
	for i, _ := range chars {
		if chars[i] == "{" {
			strack.Push("}")
			continue
		} else if chars[i] == "[" {
			strack.Push("]")
			continue
		} else if chars[i] == "(" {
			strack.Push(")")
			continue
		}
		if chars[i] != strack.Pop() {
			return false
		}
	}
	if len(strack.str) == 0 {
		return true
	}
	return false
}
