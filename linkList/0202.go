package list

//ListNode for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func kthToLast(head *ListNode, k int) int {
	node := head
	for i := 0; i < k; i++ {
		node = node.Next
	}
	for node != nil {
		node = node.Next
		head = head.Next
	}
	return head.Val
}
