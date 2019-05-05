package main

import "fmt"

//给排序好的链表去掉重复项

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	travel := head

	for travel.Next != nil {
		if travel.Next.Val == travel.Val {
			travel.Next = travel.Next.Next
		} else {
			travel = travel.Next
		}
	}
	return head

}

func main() {
	l := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3}}}
	s := deleteDuplicates(l)
	for s != nil {
		fmt.Println(s.Val)
		s = s.Next
	}
}
