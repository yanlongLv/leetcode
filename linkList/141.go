package linkList

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast, slow := head.Next, head
	for fast != nil && fast.Next != nil {
		if fast == slow {
			return true
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return false
}
