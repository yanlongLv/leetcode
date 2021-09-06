package linkList

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	node1, node2 := head, dummy
	for i := 0; i < n; i++ {
		node1 = node1.Next
	}
	for node1 != nil {
		node1 = node1.Next
		node2 = node2.Next
	}
	node2.Next = node2.Next.Next
	return dummy.Next
}
