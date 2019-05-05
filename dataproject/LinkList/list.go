package main

import (
	"dataproject/LinkList/question"
	"fmt"
)

type MyLinkedList struct {
	Val    int
	Next   *MyLinkedList
	Length int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{Next: new(MyLinkedList), Length: -1}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	temp := this.Next // 起始点
	if this.Length >= index {
		for i := 0; i <= index; i++ {
			temp = temp.Next
		}
		return temp.Val
	}
	return -1
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	temp := this.Next.Next
	head := new(MyLinkedList)
	head.Val = val
	head.Next = temp
	this.Next.Next = head
	this.Length++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	temp := this.Next
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = new(MyLinkedList)
	temp.Next.Val = val
	this.Length++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	temp := this.Next // 起始点
	if this.Length+1 >= index {
		for i := 0; i < index; i++ {
			temp = temp.Next
		}
		t := temp.Next
		node := new(MyLinkedList)
		node.Val = val
		node.Next = t
		temp.Next = node
		this.Length++
	}
	return
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if this.Length == -1 {
		return
	}
	if this.Length < index {
		return
	}
	prev := this.Next
	for i := 0; i < index; i++ {
		prev = prev.Next
	}
	if prev.Next.Next != nil {
		prev.Next = prev.Next.Next
	} else {
		prev.Next = nil
	}
	this.Length--
}

func (this *MyLinkedList) GetVals() []int {
	temp := this.Next
	result := []int{}
	for temp != nil {
		result = append(result, temp.Val)
		temp = temp.Next
	}
	return result[1:]
}

func main() {
	list := question.CreateList([]int{2, 5, 3, 4, 6, 2, 2})
	//list = question.BfsSwapPairs(list) //交换仙灵的两组数
	list = question.BfsReverseList(list) //反转链表
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}
