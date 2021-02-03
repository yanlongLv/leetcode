package list

// * Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	node := head
	length := 1
	for node.Next != nil {
		node = node.Next
		length++
	}
	node.Next = head
	k = k % length
	for i := 1; i <= length-k; i++ {
		node = node.Next
	}
	head = node.Next
	node.Next = nil
	return head
}
