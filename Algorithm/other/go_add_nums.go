package main

import (
	"fmt"
)

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//利用链表添加两个数字
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		l1 = &ListNode{}
	}
	if l2 == nil {
		l2 = &ListNode{}
	}
	result := &ListNode{}
	result.Val = l1.Val + l2.Val
	if result.Val > 9 {
		result.Val -= 10
		if l1.Next == nil {
			l1.Next = &ListNode{}
		}
		l1.Next.Val += 1
	}
	if l1.Next != nil || l2.Next != nil {
		result.Next = addTwoNumbers(l1.Next, l2.Next)
	}
	return result
}

func main() {
	l := new(ListNode)
	l.Val = 1
	l.Next = new(ListNode)
	l.Next.Val = 1
	l.Next.Next = new(ListNode)
	l.Next.Next.Val = 1
	l3 := new(ListNode)
	l3.Val = 1
	l3.Next = new(ListNode)
	l3.Next.Val = 1
	l2 := addTwoNumbers(l, l3)
	fmt.Println(l2)
}
