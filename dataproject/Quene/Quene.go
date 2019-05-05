package main

import "fmt"

type MyCircularQueue struct {
	Size int //尺寸
	Head int //头
	Tail int //尾
	node *Node
}

type Node struct {
	Val  int //值
	Next *Node
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{Size: k, node: new(Node), Head: -1, Tail: -1}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	if this.IsEmpty() {
		this.Head = 0
	}
	node1 := new(Node)
	node1.Val = value
	temp := this.node
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = node1
	this.Tail = (this.Tail + 1) % this.Size
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		this.Tail = -1
		this.Head = -1
		return false
	}
	this.node.Next = this.node.Next.Next
	if this.Head == this.Tail {
		this.Head = -1
		this.Tail = -1
		return true
	}
	this.Head = (this.Head + 1) % this.Size
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.node.Next.Val
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	temp := this.node.Next
	for temp.Next != nil {
		temp = temp.Next
	}
	return temp.Val
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	if this.Head == -1 {
		return true
	}
	return false
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	if (this.Tail+1)%this.Size == this.Head {
		return true
	}
	return false
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */

func main() {
	obj := Constructor(9)
	//添加一个元素进入循环队列
	fmt.Println(obj.EnQueue(3))
	fmt.Println(obj.Rear())
	fmt.Println(obj.EnQueue(2))
	fmt.Println(obj.Front())
	fmt.Println(obj.Front())
	fmt.Println(obj.IsEmpty())
	fmt.Println(obj.Front())
	fmt.Println(obj.EnQueue(6))
	fmt.Println(obj.EnQueue(9))
	fmt.Println(obj.DeQueue())
	fmt.Println(obj.EnQueue(5))
	fmt.Println(obj.DeQueue())
	//fmt.Println(obj.DeQueue())
	//fmt.Println(obj.DeQueue())
	//obj.node = obj.node.Next
	//for obj.node != nil {
	//	fmt.Println(obj.node.Val)
	//	obj.node = obj.node.Next
	//}

	//fmt.Println(obj.DeQueue())
	//fmt.Println(obj.EnQueue(5))
	//fmt.Println(obj.Rear())
	//判断队列是否为空
	//for i := 0;i<5;i++ {
	//	obj.EnQueue(i)
	//}
	//fmt.Println(obj.IsFull())
}
