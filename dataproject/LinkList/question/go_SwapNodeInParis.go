package question

func BfsSwapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	cur := head.Next // 当前
	prev := head     // 前一个
	flag := true
	var temp *ListNode
	if cur != nil {
		if cur.Next != nil {
			if flag {
				head = cur
				flag = false
			}
			for cur.Next.Next != nil {
				prev.Next = cur.Next.Next
				temp = cur.Next
				cur.Next = prev
				prev = temp
				cur = prev.Next
				if cur.Next == nil {
					temp = nil
					break
				}
			}
		} else {
			cur.Next = prev
			prev.Next = temp
			return cur
		}
		temp = cur.Next
		cur.Next = prev
		prev.Next = temp
	}
	return head
}

func DfsswapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	second := head.Next
	head.Next = DfsswapPairs(second.Next)
	second.Next = head
	return second
}
