package main

import "fmt"

// ========================== 栈结构
type StackInt struct {
	Number []int
	length int
}

func createStackInt(size int) StackInt {
	return StackInt{Number: make([]int, size), length: -1}
}

func (this *StackInt) Push(index int) {
	if this.length == -1 {
		this.length = 0
	} else {
		this.length++
	}
	this.Number[this.length] = index
}

func (this *StackInt) Pop() (index int) {
	if this.length >= 0 {
		index = this.Number[this.length]
		this.Number[this.length] = 0
		this.length--
	}
	return index
}

func (this *StackInt) IsEmpty() bool {
	if this.length >= 0 {
		return false
	}
	return true
}

func (this *StackInt) Top() int {
	if !this.IsEmpty() {
		return this.Number[this.length]
	}
	return 0
}

// ================================================

//每日温度
func dailyTemperatures(T []int) []int {
	length := len(T)
	if length == 0 {
		return nil
	}
	stack := createStackInt(length)
	daily := make([]int, length)
	for i := length - 1; i >= 0; i-- {
		for !stack.IsEmpty() && T[stack.Top()] <= T[i] {
			stack.Pop()
		}
		if !stack.IsEmpty() {
			fmt.Println(stack.Top(), i)
			daily[i] = stack.Top() - i
		}
		stack.Push(i)
	}
	return daily[:]
}

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}
