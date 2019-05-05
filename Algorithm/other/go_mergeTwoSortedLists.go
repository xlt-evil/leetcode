package main

import (
	"fmt"
	"sort"
)

type ListNode1 struct {
	Val  int
	Next *ListNode1
}

//* Definition for singly-linked list.
func mergeTwoLists(l1 *ListNode1, l2 *ListNode1) *ListNode1 {
	integer := []int{}
	for l1 != nil {
		integer = append(integer, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		integer = append(integer, l2.Val)
		l2 = l2.Next
	}
	sort.Ints(integer)
	Ls := []*ListNode1{}
	for i, _ := range integer {
		l := new(ListNode1)
		l.Val = integer[i]
		Ls = append(Ls, l)
	}
	var temp *ListNode1
	if len(Ls) > 0 {
		temp = Ls[0]
		for i, _ := range Ls {
			if i == 0 {
				continue
			}
			createListNode(temp, Ls[i])
		}
	}
	return temp
}

func createListNode(l1 *ListNode1, l2 *ListNode1) {
	if l1.Next != nil {
		createListNode(l1.Next, l2)
		return
	}
	l1.Next = l2
}

func main() {
	l := &ListNode1{Val: 1, Next: &ListNode1{Val: 2, Next: &ListNode1{Val: 3}}}
	s := &ListNode1{Val: 3, Next: &ListNode1{Val: 4, Next: &ListNode1{Val: 5}}}
	ss := mergeTwoLists(l, s)
	for ss != nil {
		fmt.Println(ss.Val)
		ss = ss.Next
	}
}
