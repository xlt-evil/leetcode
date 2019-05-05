package question

//
//func reverseKGroup(head *ListNode, k int) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//	stack := []*ListNode{}
//	temp := head
//	n := 1
//	for temp != nil {
//		if n%k == 0 {
//			head = temp.Next
//			n = 1
//			stack = append(stack,DfsReverseList(temp))
//			temp = head
//			continue
//		}
//		if temp.Next == nil {
//			stack = append(stack,head)
//			break
//		}
//		temp = temp.Next
//		n++
//	}
//
//	if len(stack) != 0 {
//		for _,v := range stack {
//			for v.Next!= nil {
//				v = v.Next
//			}
//			v.
//		}
//	}
//	return nil
//}

//反转链表实现
func BfsReverseList(head *ListNode) *ListNode {
	var preverseHead *ListNode // 反转的头结点
	var prev *ListNode         // 指向前一个节点
	pnode := head              // 指向头结点
	for pnode != nil {
		pnext := pnode.Next // 指向下一个元素
		if pnext == nil {   //如果next没有值就说明遍历到结尾了
			preverseHead = pnode
		}
		pnode.Next = prev //下一个指向上一个 实现反转效果
		prev = pnode      //保存当前点
		pnode = pnext
	}
	return preverseHead
}

//1 2 3 4
func DfsReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = nil
	reverseHead := DfsReverseList(next)
	next.Next = head
	return reverseHead
}
