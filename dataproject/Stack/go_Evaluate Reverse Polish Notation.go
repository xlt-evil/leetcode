package main

import (
	"fmt"
	"strconv"
)

//逆波兰表达式求值 后缀表达式

//type StackInt struct {
//	Number []int
//	length int
//}
//
//func createStackInt(size int) StackInt {
//	return StackInt{Number: make([]int, size), length: -1}
//}
//
//func (this *StackInt) Push(index int) {
//	if this.length == -1 {
//		this.length = 0
//	} else {
//		this.length++
//	}
//	this.Number[this.length] = index
//}
//
//func (this *StackInt) Pop() (index int) {
//	if this.length >= 0 {
//		index = this.Number[this.length]
//		this.Number[this.length] = 0
//		this.length--
//	}
//	return index
//}
//
//func (this *StackInt) IsEmpty() bool {
//	if this.length >= 0 {
//		return false
//	}
//	return true
//}
//
//func (this *StackInt) Top() int {
//	if !this.IsEmpty() {
//		return this.Number[this.length]
//	}
//	return 0
//}

func evalRPN(tokens []string) int {
	length := len(tokens)
	if length == 0 {
		return 0
	}
	stack := createStackInt(length)
	for i, _ := range tokens {
		switch tokens[i] {
		case "+":
			a := stack.Pop()
			b := stack.Pop()
			stack.Push(a + b)
		case "-":
			a := stack.Pop()
			b := stack.Pop()
			stack.Push(b - a)
		case "*":
			a := stack.Pop()
			b := stack.Pop()
			stack.Push(a * b)
		case "/":
			a := stack.Pop()
			b := stack.Pop()
			stack.Push(b / a)
		default:
			temp, _ := strconv.Atoi(tokens[i])
			stack.Push(temp)
		}
	}
	return stack.Pop()
}

func main() {
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}
