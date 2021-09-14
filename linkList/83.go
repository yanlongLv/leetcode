package linkList

func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dumny := &ListNode{Next: head}
	current := dumny
	for current.Next != nil && current.Next.Next != nil {
		if current.Next.Val == current.Next.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return dumny.Next
}
