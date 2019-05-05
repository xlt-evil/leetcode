package main

import "fmt"

//栈 --- 后进先出 ---- LIFO

// 最小栈 基本实现方式  设置另外一个栈来保存当前栈最小值的顺序
type MinStack struct {
	Data     []int
	MinIndex []int
	Head     int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{Data: []int{}, MinIndex: []int{}, Head: -1}
}

func (this *MinStack) Push(x int) {
	if this.Head == -1 {
		this.Head = 0
		this.Data = append(this.Data, x)
		this.MinIndex = append(this.MinIndex, this.Head)
	} else {
		this.Head++
		if this.Data[this.MinIndex[len(this.MinIndex)-1]] > x {
			this.MinIndex = append(this.MinIndex, this.Head)
		}
		this.Data = append(this.Data, x)
	}
}

func (this *MinStack) Pop() {
	if this.Head >= 0 {
		if this.Head == this.MinIndex[len(this.MinIndex)-1] {
			this.MinIndex = append(this.MinIndex[:len(this.MinIndex)-1])
		}
		this.Data = append(this.Data[:this.Head])
		this.Head--
	}
}

func (this *MinStack) Top() int {
	return this.Data[this.Head]
}

func (this *MinStack) GetMin() int {
	return this.Data[this.MinIndex[len(this.MinIndex)-1]]
}

func main() {
	stack := Constructor()
	fmt.Println(stack)
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	fmt.Println(stack.GetMin())
	stack.Pop()
	fmt.Println(stack.Top())
	fmt.Println(stack.GetMin())

}
