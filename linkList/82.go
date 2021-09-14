package linkList

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dumny := &ListNode{Next: head}
	current := dumny
	for current.Next != nil && current.Next.Next != nil {
		val := current.Next.Val
		nextVal := current.Next.Next.Val
		if val == nextVal {
			for current.Next != nil && val == current.Next.Val {
				current.Next = current.Next.Next
			}
		} else {
			current = current.Next
		}
	}
	return dumny.Next
}
