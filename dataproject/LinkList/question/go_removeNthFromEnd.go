package question

type ListNode struct {
	Val  int
	Next *ListNode
}

//移除链表从后往前的元素
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	stack := []*ListNode{}
	temp := head
	for temp != nil {
		//加入栈
		stack = append(stack, temp)
		temp = temp.Next
	}
	length := len(stack)
	if length < n {
		return head
	}
	//删除头元素
	if length-n-1 < 0 {
		if length > 1 {
			return stack[1]
		}
		return nil
	}
	prev := stack[length-n-1]
	prev.Next = stack[length-n].Next
	return stack[0]
}

func CreateList(result []int) *ListNode {
	if len(result) == 0 {
		return nil
	}
	head := new(ListNode)
	head.Val = result[0]
	temp := head
	for _, v := range result[1:] {
		s := new(ListNode)
		s.Val = v
		temp.Next = s
		temp = s
	}
	return head
}
