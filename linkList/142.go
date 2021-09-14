package linkList

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}

	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}

	return slow
}

// func detectCycle(head *ListNode) *ListNode {
// 	if head == nil {
// 		return head
// 	}
// 	fast, slow := head, head
// 	for fast != nil && fast.Next != nil {
// 		fast = fast.Next.Next
// 		slow = slow.Next
// 		if fast == slow {
// 			break
// 		}
// 	}
// 	fast = head
// 	for fast != slow {
// 		fast = fast.Next
// 		slow = slow.Next
// 	}
// 	return slow
// }
